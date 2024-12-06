package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

// ComponentInfo holds metadata about a component
type ComponentInfo struct {
	Name       string
	Package    string
	Fields     []FieldInfo
	Imports    []string
	DefaultVar string // Name of the default variable if present
}

// FieldInfo represents a component field with its metadata
type FieldInfo struct {
	Name         string
	Type         string
	PackageName  string // Package name for imported types
	IsPointer    bool
	IsSlice      bool
	DefaultValue string
	ImportPath   string
	Comment      string
}

// parseComponent analyzes a Go file to extract component information
func parseComponent(filename string) (*ComponentInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parsing file: %w", err)
	}

	comp := &ComponentInfo{
		Name:    filepath.Base(filepath.Dir(filename)),
		Package: node.Name.Name,
	}

	// Extract imports and create import map
	imports := make(map[string]string)
	for _, imp := range node.Imports {
		path := strings.Trim(imp.Path.Value, `"`)
		name := filepath.Base(path)
		if imp.Name != nil {
			name = imp.Name.Name
		}
		imports[name] = path
		comp.Imports = append(comp.Imports, path)
	}

	// Look for DEFAULTS variable or init function
	ast.Inspect(node, func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.ValueSpec:
			if len(v.Names) > 0 && v.Names[0].Name == "DEFAULTS" {
				comp.DefaultVar = "DEFAULTS"
			}
		case *ast.FuncDecl:
			if v.Name.Name == "init" {
				// Found init function, component might have defaults
				comp.DefaultVar = "DEFAULTS"
			}
		}
		return true
	})

	// Find and parse the D struct
	ast.Inspect(node, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok || typeSpec.Name.Name != "D" {
			return true
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		// Parse struct fields
		for _, field := range structType.Fields.List {
			if field.Names == nil {
				continue
			}

			fieldInfo := parseField(field, imports)
			comp.Fields = append(comp.Fields, fieldInfo)
		}
		return false
	})

	return comp, nil
}

// parseField extracts field information
func parseField(field *ast.Field, imports map[string]string) FieldInfo {
	info := FieldInfo{
		Name:    field.Names[0].Name,
		Comment: extractComment(field.Doc),
	}

	// Parse type information
	switch t := field.Type.(type) {
	case *ast.ArrayType:
		info.IsSlice = true
		info = parseTypeInfo(t.Elt, info, imports)
	case *ast.StarExpr:
		info.IsPointer = true
		info = parseTypeInfo(t.X, info, imports)
	default:
		info = parseTypeInfo(field.Type, info, imports)
	}

	return info
}

// parseTypeInfo processes type information and updates FieldInfo
func parseTypeInfo(expr ast.Expr, info FieldInfo, imports map[string]string) FieldInfo {
	switch t := expr.(type) {
	case *ast.Ident:
		info.Type = t.Name
	case *ast.SelectorExpr:
		pkgName := t.X.(*ast.Ident).Name
		info.Type = fmt.Sprintf("%s.%s", pkgName, t.Sel.Name)
		info.PackageName = pkgName
		if imp, ok := imports[pkgName]; ok {
			info.ImportPath = imp
		}
	}
	return info
}

// Story generation template
const storyTemplate = `
package {{.Package}}

import (
    "github.com/joaorufino/templui/internal/storybook"
    {{range .Imports}}
    "{{.}}"
    {{end}}
)

func Add{{.Name}}Story(s *storybook.Storybook) {

    s.AddComponent("{{.Name}}",
        func({{range .Fields}}
            {{safeVariableName .Name}} {{if .IsSlice}}[]{{end}}{{.Type}},{{end}}
        ) templ.Component {
            def := D{
                {{range .Fields}}
                {{.Name}}: {{safeVariableName .Name}},
                {{end}}
            }
            return C(def)
        },
        {{range .Fields}}
        {{generateArg . $.DefaultVar}},
        {{end}}
    )
}
`

// generateArg determines the appropriate argument type for a field
func generateArg(field FieldInfo, defaultVar string) string {
	// Helper to create default value reference
	getDefault := func(field FieldInfo) string {
		if defaultVar != "" {
			return fmt.Sprintf("DEFAULTS.%s", field.Name)
		}
		return getEmptyValue(field)
	}

	safeName := safeVariableName(field.Name)
	switch {
	case field.Type == "string":
		return fmt.Sprintf(`storybook.TextArg("%s", %s)`,
			safeName,
			getDefault(field))

	case field.Type == "bool":
		return fmt.Sprintf(`storybook.BooleanArg("%s", %s)`,
			safeName,
			getDefault(field))

	case field.Type == "int":
		return fmt.Sprintf(`storybook.IntArg("%s", %s)`,
			safeName,
			getDefault(field))

	case field.Type == "map[string]string":
		return fmt.Sprintf(`storybook.MapArg("%s", %s)`,
			safeName,
			getDefault(field))

	case field.IsSlice:
		return fmt.Sprintf(`storybook.SliceArg("%s", %s)`,
			safeName,
			getDefault(field))

	case strings.HasPrefix(field.Type, "style.Custom"):
		return fmt.Sprintf(`storybook.CustomStyleArg("%s", %s)`,
			safeName,
			getDefault(field))

	case strings.HasPrefix(field.Type, "style.Style"):
		return fmt.Sprintf(`storybook.StyleArg("%s", %s)`,
			safeName,
			getDefault(field))

	case strings.HasSuffix(field.Type, ".D"):
		return fmt.Sprintf(`storybook.ComponentArg("%s", %s.DEFAULTS)`,
			safeName,
			field.PackageName)

	default:
		return fmt.Sprintf(`storybook.ObjectArg("%s", %s)`,
			safeName,
			getDefault(field))
	}
}

// getEmptyValue returns appropriate empty/zero value for a type
func getEmptyValue(field FieldInfo) string {
	switch {
	case field.Type == "string":
		return `""`
	case field.Type == "bool":
		return "false"
	case field.Type == "int":
		return "0"
	case field.IsSlice:
		return fmt.Sprintf("make([]%s, 0)", field.Type)
	case field.IsPointer:
		return "nil"
	case strings.HasSuffix(field.Type, ".D"):
		return fmt.Sprintf("%s.D{}", field.PackageName)
	default:
		return fmt.Sprintf("%s{}", field.Type)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: storygen <components-dir>")
	}

	tmpl := template.Must(template.New("story").Funcs(template.FuncMap{
		"lower":            strings.ToLower,
		"generateArg":      generateArg,
		"safeVariableName": safeVariableName,
	}).Parse(storyTemplate))

	// Process components
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".go") ||
			strings.HasSuffix(path, "_test.go") ||
			strings.HasSuffix(path, "_templ.go") {
			return nil
		}

		comp, err := parseComponent(path)
		if err != nil {
			log.Printf("Warning: Skipping %s: %v", path, err)
			return nil
		}

		if len(comp.Fields) == 0 {
			return nil
		}

		// Generate story file
		outFile := filepath.Join(filepath.Dir(path), fmt.Sprintf("%s_story.go", comp.Name))
		f, err := os.Create(outFile)
		if err != nil {
			return fmt.Errorf("creating output file: %w", err)
		}
		defer f.Close()

		if err := tmpl.Execute(f, comp); err != nil {
			return fmt.Errorf("executing template: %w", err)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

// commentProcessor holds state and configuration for comment processing
type commentProcessor struct {
	keepTags        bool // Whether to preserve documentation tags like @param
	stripDirectives bool // Whether to remove go:generate and other directives
	maxLength       int  // Maximum length for processed comments
	unwrapText      bool // Whether to unwrap multi-line text into single lines
}

// extractComment processes an AST comment group into a clean, usable comment string.
// It handles various comment formats and documentation styles commonly found in Go code.
func extractComment(doc *ast.CommentGroup) string {
	if doc == nil {
		return ""
	}

	processor := commentProcessor{
		keepTags:        true, // Preserve documentation tags by default
		stripDirectives: true, // Remove Go directives
		maxLength:       1000, // Reasonable limit for comment length
		unwrapText:      true, // Combine wrapped lines
	}

	// Process each comment line
	var lines []string
	for _, comment := range doc.List {
		line := processor.processCommentLine(comment.Text)
		if line != "" {
			lines = append(lines, line)
		}
	}

	// Combine and clean up the final result
	result := processor.combineLines(lines)
	return processor.finalCleanup(result)
}

// processCommentLine handles a single comment line, removing markers and extra whitespace
func (p *commentProcessor) processCommentLine(text string) string {
	// Remove comment markers
	text = strings.TrimPrefix(text, "//")
	text = strings.TrimPrefix(text, "/*")
	text = strings.TrimSuffix(text, "*/")

	// Skip directives if configured
	if p.stripDirectives {
		if strings.HasPrefix(strings.TrimSpace(text), "+") ||
			strings.HasPrefix(strings.TrimSpace(text), "go:") {
			return ""
		}
	}

	// Clean up the text
	text = strings.TrimSpace(text)

	// Handle playground tags specially
	if strings.Contains(text, "playground:") {
		return p.processPlaygroundTag(text)
	}

	return text
}

// processPlaygroundTag extracts useful information from playground tags
func (p *commentProcessor) processPlaygroundTag(text string) string {
	// Extract values from playground tags
	if strings.Contains(text, "playground:values:") {
		values := strings.TrimPrefix(text, "playground:values:")
		return fmt.Sprintf("default values: %s", values)
	}
	// Handle other playground tags if needed
	return ""
}

// combineLines joins multiple comment lines into a single coherent text
func (p *commentProcessor) combineLines(lines []string) string {
	if len(lines) == 0 {
		return ""
	}

	// If we don't want to unwrap text, just join with newlines
	if !p.unwrapText {
		return strings.Join(lines, "\n")
	}

	// Combine wrapped lines intelligently
	var result strings.Builder
	var paragraph []string

	for _, line := range lines {
		// Check for paragraph breaks
		if line == "" {
			if len(paragraph) > 0 {
				result.WriteString(strings.Join(paragraph, " "))
				result.WriteString("\n\n")
				paragraph = paragraph[:0]
			}
			continue
		}

		// Check if this line starts a new sentence or continues the previous one
		if len(paragraph) > 0 && !strings.Contains(".-*•", string(line[0])) && !unicode.IsUpper(rune(line[0])) {
			// Continue previous paragraph
			paragraph = append(paragraph, line)
		} else {
			// Start new paragraph
			if len(paragraph) > 0 {
				result.WriteString(strings.Join(paragraph, " "))
				result.WriteString("\n\n")
			}
			paragraph = []string{line}
		}
	}

	// Add final paragraph
	if len(paragraph) > 0 {
		result.WriteString(strings.Join(paragraph, " "))
	}

	return strings.TrimSpace(result.String())
}

// finalCleanup performs final processing on the complete comment text
func (p *commentProcessor) finalCleanup(text string) string {
	// Remove consecutive blank lines
	text = regexp.MustCompile(`\n{3,}`).ReplaceAllString(text, "\n\n")

	// Truncate if needed
	if p.maxLength > 0 && len(text) > p.maxLength {
		return text[:p.maxLength-3] + "..."
	}

	return text
}

func safeVariableName(fieldName string) string {
	// Map of Go reserved words to safe alternatives
	reservedWords := map[string]string{
		"type":      "typeValue",
		"range":     "rangeValue",
		"select":    "selectValue",
		"func":      "funcValue",
		"defer":     "deferValue",
		"go":        "goValue",
		"map":       "mapValue",
		"chan":      "chanValue",
		"package":   "packageValue",
		"interface": "interfaceValue",
		"switch":    "switchValue",
		"case":      "caseValue",
		"default":   "defaultValue",
		"break":     "breakValue",
		"continue":  "continueValue",
		"for":       "forValue",
		"if":        "ifValue",
		"else":      "elseValue",
		"var":       "varValue",
		"const":     "constValue",
		"import":    "importValue",
		"return":    "returnValue",
		"struct":    "structValue",
		"style":     "styleValue",
		"size":      "sizeValue",
	}

	// Convert to lower case for consistency
	name := strings.ToLower(fieldName)

	// Check if it's a reserved word
	if replacement, isReserved := reservedWords[name]; isReserved {
		return replacement
	}

	return name
}
