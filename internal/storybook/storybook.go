package storybook

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joaorufino/templui/components/style"

	"github.com/a-h/pathvars"
	"github.com/a-h/templ"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Storybook struct {
	Path               string
	Config             map[string]*Conf
	Handlers           map[string]http.Handler
	StaticHandler      http.Handler
	Server             http.Server
	Log                *zap.Logger
	AdditionalPrefixJS string
}

type StorybookConfig func(*Storybook)

func WithServerAddr(addr string) StorybookConfig {
	return func(sb *Storybook) {
		sb.Server.Addr = addr
	}
}

func New(conf ...StorybookConfig) *Storybook {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, err := cfg.Build()
	if err != nil {
		panic("templ-storybook: zap configuration failed: " + err.Error())
	}
	sh := &Storybook{
		Path:     ".",
		Config:   map[string]*Conf{},
		Handlers: map[string]http.Handler{},
		Log:      logger,
	}
	sh.StaticHandler = http.FileServer(http.Dir(path.Join(sh.Path, "storybook-static")))
	sh.Server = http.Server{
		Handler:           sh,
		Addr:              ":60606",
		ReadHeaderTimeout: 2 * time.Second,
	}
	for _, sc := range conf {
		sc(sh)
	}
	return sh
}

func (sh *Storybook) AddComponent(name string, componentConstructor interface{}, args ...Arg) {
	c := NewConf(name, args...)
	sh.Config[name] = c
	h := NewHandler(name, componentConstructor, args...)
	sh.Handlers[name] = h
}

var storybookPreviewMatcher = pathvars.NewExtractor("/storybook_preview/{name}")

func (sh *Storybook) Build(ctx context.Context) error {
	defer sh.Log.Sync()

	sh.Log.Info("Installing Storybook.")
	if err := sh.installStorybook(); err != nil {
		return err
	}
	// Then install the addons
	if err := sh.installAddons(); err != nil {
		return err
	}

	sh.Log.Info("Configuring Storybook.")
	configHasChanged, err := sh.configureStorybook()
	if err != nil {
		return err
	}

	if configHasChanged {
		sh.Log.Info("Building Storybook.")
		if err := sh.buildStorybook(); err != nil {
			return err
		}
	} else {
		sh.Log.Info("Storybook is up-to-date, skipping build.")
	}

	return nil
}

func (sh *Storybook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/debug/") {
		sh.debugHandler(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/storybook_preview/") {
		sh.previewHandler(w, r)
		return
	}
	sh.StaticHandler.ServeHTTP(w, r)
}

func (sh *Storybook) ListenAndServeWithContext(ctx context.Context) error {
	if err := sh.Build(ctx); err != nil {
		return err
	}

	go func() {
		sh.Log.Info("Starting server", zap.String("address", sh.Server.Addr))
		if err := sh.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sh.Log.Error("Server error", zap.Error(err))
		}
	}()

	<-ctx.Done()
	sh.Server.Close()
	return nil
}

func (sh *Storybook) previewHandler(w http.ResponseWriter, r *http.Request) {
	values, ok := storybookPreviewMatcher.Extract(r.URL)
	if !ok {
		http.NotFound(w, r)
		return
	}
	name := values["name"]
	h, found := sh.Handlers[name]
	if !found {
		http.NotFound(w, r)
		return
	}
	h.ServeHTTP(w, r)
}

func (sh *Storybook) installStorybook() error {
	if _, err := os.Stat(filepath.Join(sh.Path, "package.json")); err == nil {
		sh.Log.Info("Storybook already installed, skipping.")
		return nil
	}

	cmd := exec.Command("pnpm", "dlx", "storybook@latest", "init", "-t", "server", "--builder", "webpack5")
	cmd.Dir = sh.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error initializing Storybook: %w", err)
	}
	return nil
}

// installAddons handles the installation of Storybook addons
func (sh *Storybook) installAddons() error {
	if _, err := os.Stat(filepath.Join(sh.Path, "package.json")); err == nil {
		sh.Log.Info("Storybook addons already installed, skipping.")
		return nil
	}
	// Define commonly used addons
	addons := []string{
		"@storybook/addon-essentials",
		"@storybook/addon-controls",
		"@storybook/addon-actions",
		"@storybook/addon-docs",
		"@storybook/addon-links",
	}

	// Install each addon using pnpm
	for _, addon := range addons {
		sh.Log.Info("Installing addon: " + addon)
		cmd := exec.Command("pnpm", "add", "-D", addon)
		cmd.Dir = sh.Path
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error installing addon %s: %w", addon, err)
		}
	}

	return nil
}

func (sh *Storybook) configureStorybook() (bool, error) {
	storiesDir := filepath.Join(sh.Path, "stories")
	beforeHash, _ := hashDirectory(storiesDir)
	os.RemoveAll(storiesDir)
	os.MkdirAll(storiesDir, os.ModePerm)

	for _, c := range sh.Config {
		name := filepath.Join(sh.Path, fmt.Sprintf("stories/%s.stories.json", c.Title))
		f, err := os.Create(name)
		if err != nil {
			return false, fmt.Errorf("failed to create config file %q: %w", name, err)
		}
		err = json.NewEncoder(f).Encode(c)
		f.Close()
		if err != nil {
			return false, fmt.Errorf("failed to write JSON config to %q: %w", name, err)
		}
	}

	afterHash, _ := hashDirectory(storiesDir)
	configHasChanged := beforeHash != afterHash

	err := os.WriteFile(filepath.Join(sh.Path, ".storybook/preview.js"), []byte(fmt.Sprintf("%s\n%s", sh.AdditionalPrefixJS, previewJS)), os.ModePerm)
	if err != nil {
		return configHasChanged, err
	}
	return configHasChanged, nil
}

var previewJS = `
export const parameters = {
  server: {
    url: window.location.origin + "/storybook_preview",
  },
};
`

func (sh *Storybook) buildStorybook() error {
	cmd := exec.Command("npm", "run", "build-storybook")
	cmd.Dir = sh.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// debugHandler prints the current state of the Storybook
func (sh *Storybook) debugHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the content type is JSON
	w.Header().Set("Content-Type", "application/json")

	// Serialize the Storybook struct to JSON
	state := struct {
		Path     string
		Config   map[string]*Conf
		Handlers []string
	}{
		Path:   sh.Path,
		Config: sh.Config,
	}

	// Extract handler names for readability
	for key := range sh.Handlers {
		state.Handlers = append(state.Handlers, key)
	}

	// Serialize the state to JSON
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		http.Error(w, "Failed to serialize state", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(data)
}

func NewHandler(name string, f interface{}, args ...Arg) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		argv := make([]interface{}, len(args))
		q := r.URL.Query()
		for i, arg := range args {
			argv[i] = arg.Get(q)
		}
		component, err := executeTemplate(name, f, argv)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		templ.Handler(component).ServeHTTP(w, r)
	})
}

func executeTemplate(name string, fn interface{}, values []interface{}) (templ.Component, error) {
	v := reflect.ValueOf(fn)
	t := v.Type()

	if t.Kind() != reflect.Func {
		return nil, fmt.Errorf("%s is not a function", name)
	}
	if t.NumIn() != len(values) {
		return nil, fmt.Errorf("function %s expects %d arguments, got %d", name, t.NumIn(), len(values))
	}

	argv := make([]reflect.Value, len(values))
	for i, val := range values {
		paramType := t.In(i)
		valReflect := reflect.ValueOf(val)

		// Check if the parameter expects a pointer type
		if paramType.Kind() == reflect.Ptr {
			// If the value is already a pointer, ensure it matches the expected type
			if valReflect.Type() == paramType {
				argv[i] = valReflect // Use as-is if it matches the pointer type
			} else if valReflect.Type() == paramType.Elem() {
				// If it's the element type, take its address
				if !valReflect.CanAddr() {
					return nil, fmt.Errorf("value for parameter %d cannot be taken as a pointer", i)
				}
				argv[i] = valReflect.Addr()
			} else {
				return nil, fmt.Errorf("value for parameter %d is of type %s; expected %s or %s", i, valReflect.Type(), paramType, paramType.Elem())
			}
		} else {
			// Ensure value matches expected non-pointer type
			if valReflect.Type() != paramType {
				return nil, fmt.Errorf("value for parameter %d is of type %s; expected %s", i, valReflect.Type(), paramType)
			}
			argv[i] = valReflect
		}
	}

	results := v.Call(argv)
	if len(results) != 1 {
		return nil, fmt.Errorf("function %s must return exactly one result", name)
	}

	component, ok := results[0].Interface().(templ.Component)
	if !ok {
		return nil, fmt.Errorf("result of %s is not a templ.Component", name)
	}

	return component, nil
}

func NewConf(title string, args ...Arg) *Conf {
	c := &Conf{
		Title: title,
		Parameters: StoryParameters{
			Server: map[string]interface{}{
				"id": title,
			},
		},
		Args:     NewSortedMap(),
		ArgTypes: NewSortedMap(),
		Stories:  []Story{},
	}
	for _, arg := range args {
		c.Args.Add(arg.Name, arg.Value)
		// Create a control configuration that includes options
		controlConfig := map[string]interface{}{
			"control": arg.Control,
		}

		// If we have options, add them to the control configuration
		if arg.Options != nil && len(arg.Options) > 0 {
			controlConfig["options"] = arg.Options
		}
		c.ArgTypes.Add(arg.Name, controlConfig)
	}
	c.AddStory("Default")
	return c
}

func (c *Conf) AddStory(name string, args ...Arg) {
	m := NewSortedMap()
	for _, arg := range args {
		m.Add(arg.Name, arg.Value)
	}
	c.Stories = append(c.Stories, Story{
		Name: name,
		Args: m,
	})
}

type Arg struct {
	Name    string
	Value   interface{}
	Control interface{}
	Options []interface{}
	Get     func(q url.Values) interface{}
}

func TextArg(name, value string) Arg {
	return Arg{
		Name:    name,
		Value:   value,
		Control: "text",
		Get: func(q url.Values) interface{} {
			return q.Get(name)
		},
	}
}

func BooleanArg(name string, value bool) Arg {
	return Arg{
		Name:    name,
		Value:   value,
		Control: "boolean",
		Get: func(q url.Values) interface{} {
			return q.Get(name) == "true"
		},
	}
}

func IntArg(name string, value int) Arg {
	return Arg{
		Name:    name,
		Value:   value,
		Control: "number",
		Get: func(q url.Values) interface{} {
			i, _ := strconv.Atoi(q.Get(name))
			return i
		},
	}
}

func FloatArg(name string, value float64) Arg {
	return Arg{
		Name:    name,
		Value:   value,
		Control: "number",
		Get: func(q url.Values) interface{} {
			f, _ := strconv.ParseFloat(q.Get(name), 64)
			return f
		},
	}
}

type Conf struct {
	Title      string          `json:"title"`
	Parameters StoryParameters `json:"parameters"`
	Args       *SortedMap      `json:"args"`
	ArgTypes   *SortedMap      `json:"argTypes"`
	Stories    []Story         `json:"stories"`
}

type StoryParameters struct {
	Server map[string]interface{} `json:"server"`
}

func NewSortedMap() *SortedMap {
	return &SortedMap{
		m:        new(sync.Mutex),
		internal: map[string]interface{}{},
		keys:     []string{},
	}
}

type SortedMap struct {
	m        *sync.Mutex
	internal map[string]interface{}
	keys     []string
}

func (sm *SortedMap) Add(key string, value interface{}) {
	sm.m.Lock()
	defer sm.m.Unlock()
	sm.keys = append(sm.keys, key)
	sm.internal[key] = value
}

func (sm *SortedMap) MarshalJSON() ([]byte, error) {
	sm.m.Lock()
	defer sm.m.Unlock()
	b := new(strings.Builder)
	b.WriteRune('{')
	for i, k := range sm.keys {
		jsonKey, _ := json.Marshal(k)
		jsonValue, _ := json.Marshal(sm.internal[k])
		b.Write(jsonKey)
		b.WriteRune(':')
		b.Write(jsonValue)
		if i < len(sm.keys)-1 {
			b.WriteRune(',')
		}
	}
	b.WriteRune('}')
	return []byte(b.String()), nil
}

type Story struct {
	Name string     `json:"name"`
	Args *SortedMap `json:"args"`
}

func hashDirectory(dir string) (string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	hash := ""
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return "", err
		}
		hash += fmt.Sprintf("%x", data)
	}
	return hash, nil
}

// OptionValue represents a single option in a select/dropdown control
type OptionValue struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

// OptionsArg creates a select/dropdown control with predefined options
func OptionsArg(name string, defaultValue interface{}, options []OptionValue) Arg {
	// Extract option values and labels for the argType
	optionValues := make([]interface{}, len(options))
	optionLabels := make(map[string]string)
	optionMapping := make(map[string]interface{})

	for i, opt := range options {
		optionValues[i] = opt.Value
		if opt.Label != "" {
			strValue := fmt.Sprint(opt.Value)
			optionLabels[strValue] = opt.Label
			optionMapping[strValue] = opt.Value
		}
	}

	return Arg{
		Name:    name,
		Value:   defaultValue,
		Control: "select",
		Options: optionValues, // Options array at argType level
		Get: func(q url.Values) interface{} {
			selected := q.Get(name)
			if mapping, ok := optionMapping[selected]; ok {
				return mapping
			}
			return defaultValue
		},
	}
}

func CustomArg(name string, value interface{}) Arg {
	return Arg{
		Name:    name,
		Value:   value,
		Control: "object",
		Get: func(q url.Values) interface{} {
			return value
		},
	}
}

// getStyleOptions returns available style options for a component's style field.
// It examines the type's constants to build selectable options for the storybook UI.
func getStyleOptions() []interface{} {
	// Start with common styles that apply to all components
	options := []interface{}{
		style.Default,
		style.Disabled,
		style.Valid,
		style.Invalid,
	}

	// Add size-related styles
	for _, size := range []style.Style{
		style.SizeXS,
		style.SizeS,
		style.SizeNormal,
		style.SizeL,
		style.SizeXL,
		style.SizeTwoXL,
		style.SizeThreeXL,
		style.SizeFourXL,
		style.SizeFiveXL,
		style.SizeSixXL,
		style.SizeSevenXL,
		style.SizeEightXL,
		style.SizeNineXL,
		style.SizeFull,
	} {
		options = append(options, size)
	}

	return options
}

func parseCustomStyle(value string) style.Custom {
	return style.Custom{}
}

// parseStyle converts a string representation of style(s) into style.Style.
// It handles both single styles and combinations (comma-separated).
func parseStyle(value string) style.Style {
	// Handle empty case
	if value == "" {
		return style.Default
	}

	// Split multi-style string
	parts := strings.Split(value, ",")
	var result style.Style

	// Map of style names to values
	styleMap := map[string]style.Style{
		"Default":  style.Default,
		"Disabled": style.Disabled,
		"Valid":    style.Valid,
		"Invalid":  style.Invalid,
		"SizeXS":   style.SizeXS,
		"SizeS":    style.SizeS,
		"Normal":   style.SizeNormal,
		"SizeL":    style.SizeL,
		"SizeXL":   style.SizeXL,
		"TwoXL":    style.SizeTwoXL,
		"ThreeXL":  style.SizeThreeXL,
		"FourXL":   style.SizeFourXL,
		"FiveXL":   style.SizeFiveXL,
		"SixXL":    style.SizeSixXL,
		"SevenXL":  style.SizeSevenXL,
		"EightXL":  style.SizeEightXL,
		"NineXL":   style.SizeNineXL,
		"Full":     style.SizeFull,
	}

	// Combine styles using bitwise OR
	for _, part := range parts {
		if style, ok := styleMap[strings.TrimSpace(part)]; ok {
			result |= style
		}
	}

	return result
}

// StyleArg creates an argument for style.Style fields.
// It provides a select control with available style options.
func StyleArg(name string, value style.Style) Arg {
	return Arg{
		Name:    name,
		Value:   value,
		Control: "select",
		Options: getStyleOptions(),
		Get: func(q url.Values) interface{} {
			return parseStyle(q.Get(name))
		},
	}
}

func CustomStyleArg(name string, value style.Custom) Arg {
	return Arg{
		Name:    name,
		Value:   value,
		Control: "select",
		Options: getStyleOptions(),
		Get: func(q url.Values) interface{} {
			return parseCustomStyle(q.Get(name))
		},
	}
}

// ComponentArg creates an argument for component type fields (ending in .D).
// It handles both struct fields and pointers to structs.
func ComponentArg(name string, value interface{}) Arg {
	return Arg{
		Name:    name,
		Value:   value,
		Control: "object",
		Get: func(q url.Values) interface{} {
			return parseComponent(q.Get(name), value)
		},
	}
}

// parseComponent converts a JSON string into a component instance.
// It handles both normal structs and pointers to structs.
func parseComponent(jsonStr string, template interface{}) interface{} {
	// Handle empty case
	if jsonStr == "" {
		return template
	}

	// Create new instance of the same type as template
	val := reflect.New(reflect.TypeOf(template)).Interface()

	// Attempt to parse JSON
	if err := json.Unmarshal([]byte(jsonStr), val); err != nil {
		// Return template on error
		return template
	}

	// If template is not a pointer but val is, dereference val
	if reflect.TypeOf(template).Kind() != reflect.Ptr {
		return reflect.ValueOf(val).Elem().Interface()
	}

	return val
}

// SliceArg creates an argument for slice fields.
// It provides array controls in the storybook UI.
func SliceArg[T any](name string, value []T) Arg {
	return Arg{
		Name:    name,
		Value:   value,
		Control: "array",
		Get: func(q url.Values) interface{} {
			return parseSlice(q.Get(name), value)
		},
	}
}

// parseSlice converts a JSON array string into a slice of the appropriate type.
// It preserves the element type of the template slice.
func parseSlice[T any](jsonStr string, template []T) []T {
	// Handle empty case
	if jsonStr == "" {
		return template
	}

	// Create new slice of the same type
	var result []T

	// Attempt to parse JSON
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		// Return template on error
		return template
	}

	return result
}

// Helper function to convert string to number for numeric types
func stringToNumber(s string, kind reflect.Kind) interface{} {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v, err := strconv.ParseInt(s, 10, 64); err == nil {
			return v
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v, err := strconv.ParseUint(s, 10, 64); err == nil {
			return v
		}
	case reflect.Float32, reflect.Float64:
		if v, err := strconv.ParseFloat(s, 64); err == nil {
			return v
		}
	}
	return 0
}

// ObjectArg creates an argument for general object types.
// It provides appropriate controls based on the object's structure.
func ObjectArg(name string, value interface{}) Arg {
	return Arg{
		Name:    name,
		Value:   value,
		Control: "object",
		Get: func(q url.Values) interface{} {
			jsonStr := q.Get(name)
			if jsonStr == "" {
				return value
			}

			// Create new instance of same type
			val := reflect.New(reflect.TypeOf(value)).Interface()
			if err := json.Unmarshal([]byte(jsonStr), val); err == nil {
				// Dereference if needed
				if reflect.TypeOf(value).Kind() != reflect.Ptr {
					return reflect.ValueOf(val).Elem().Interface()
				}
				return val
			}
			return value
		},
	}
}

// WithAdditionalPreviewJS / WithAdditionalPreviewJS allows to add content to the generated .storybook/preview.js file.
// For example this can be used to include custom CSS.
func (sb *Storybook) WithAdditionalPreviewJS(content string) {
	sb.AdditionalPrefixJS = content
}
