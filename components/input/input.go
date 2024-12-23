// Package input implements input fields.
package input

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/components/form/validation/message"
	"github.com/joaorufino/templui/components/label"
	"github.com/joaorufino/templui/components/position"
	"github.com/joaorufino/templui/components/size"
	"github.com/joaorufino/templui/components/style"
	"github.com/joaorufino/templui/components/tooltip"
)

var DEFAULTS = D{}

type Type string

const (
	TypeEmail    Type = "email"
	TypeNumber   Type = "number"
	TypePassword Type = "password"
	TypeText     Type = "text"
	TypeURL      Type = "url"
	TypeHidden   Type = "hidden"
)

func init() {
	style.SetDefaults(style.Defaults{
		"input/input": {
			style.Default: {
				style.Set("block w-full border rounded-lg"),
			},
			style.Disabled: {
				style.Add("cursor-not-allowed"),
			},
			style.SizeS: {
				style.Add("p-2 text-xs"),
			},
			style.SizeNormal: {
				style.Add("p-2.5 text-sm"),
			},
			style.SizeL: {
				style.Add("p-4 text-base"),
			},
		},
		"input/icon": {
			style.Default: {
				style.Set("absolute inset-y-0 flex items-center pointer-events-none"),
			},
		},
	})
}

// D is the definition for input fields.
type D struct {
	// ID is the input id (Name if not set).
	ID string
	// Name is the input name.
	Name string
	// Type is the input type (text, password, ...).
	Type Type
	// Style defines the style (style.Default, Valid or Invalid).
	Style style.Style
	// Label is the input label.
	Label any
	// Value is the input value.
	Value string
	// Placeholder is the placeholder text displayed when no value is set.
	Placeholder string
	// Message adds a validation message below the field.
	// Just add &message.D{} to add automatic validation.
	//playground:import:github.com/joaorufino/templui/components/form/validation/message
	//playground:default:&message.D{Message: "Validation message"}
	Message *message.D
	// Disabled disables the input.
	Disabled bool
	// Invalid marks the input as invalid.
	Invalid bool
	// Size defines the input size (size.S, size.Normal (default) or size.L).
	//playground:values:size.S,size.Normal,size.L
	Size size.Size
	// Loader displays a spinning loader when an HTMX action is triggered by the input.
	Loader bool
	// Icon displays an icon on the left side.
	//playground:values:icon.Flower
	Icon string
	// IconPosition can be position.Start (default) or position.End.
	//playground:values:position.Start,position.End
	IconPosition position.Position
	// Tooltip adds a tooltip to the input.
	Tooltip *tooltip.D
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"input":       style.D{style.Add("...")},
	// 		"input/input": style.D{style.Add("...")},
	// 		"input/icon":  style.D{style.Add("...")},
	// 		"input/label": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) size() size.Size {
	if def.Size == 0 {
		return size.Normal
	}
	return def.Size
}

func (def D) style() style.Style {
	st := def.Style | style.Size(def.size())
	if def.Invalid {
		return st | style.Invalid
	}
	if def.Disabled {
		return st | style.Disabled
	}
	return st
}

func (def D) label() label.D {
	cc := style.Custom{
		"label": style.Compute(def.Style, "input/label", def.CustomStyle),
	}
	switch l := def.Label.(type) {
	case string:
		return label.D{
			InputID:      def.id(),
			InputName:    def.Name,
			Label:        l,
			Style:        def.style(),
			NoValidation: def.Message == nil,
			CustomStyle:  cc,
		}
	case templ.Component:
		return label.D{
			InputID:      def.id(),
			InputName:    def.Name,
			Label:        l,
			Style:        def.style(),
			NoValidation: def.Message == nil,
			CustomStyle:  cc,
		}
	case label.D:
		l.InputID = def.id()
		l.InputName = def.Name
		if l.Style == style.Default {
			l.Style = def.style()
		}
		l.NoValidation = def.Message == nil
		l.CustomStyle = l.CustomStyle.AddBefore(cc)
		return l
	default:
		return label.D{}
	}
}

func (def D) id() string {
	if def.ID != "" {
		return def.ID
	}
	return def.Name
}

func (def D) inputType() string {
	if def.Type == "" {
		return string(TypeText)
	}
	return string(def.Type)
}

func (def D) iconClass() string {
	class := style.CSSClass(def.style(), "input/icon", def.CustomStyle)
	switch {
	case def.Icon == "":
		return ""
	case def.IconPosition == position.End:
		return class + " end-0 pe-3.5"
	default:
		return class + " start-0 ps-3.5"
	}
}

func (def D) iconSize() size.Size {
	if def.Size == size.Inherit {
		return size.Normal
	}
	return def.Size
}

func (def D) inputClass() string {
	class := style.CSSClass(def.style(), "input/input", def.CustomStyle)
	if def.Tooltip != nil {
		class += " " + def.Tooltip.Class()
	}
	switch {
	case def.Icon == "":
		return class
	case def.IconPosition == position.End:
		return class + " pe-10"
	default:
		return class + " ps-10"
	}
}

func (def D) containerClass() string {
	return style.CSSClass(def.style(), "input", def.CustomStyle)
}

func (def D) inputClassInvalid() string {
	return style.Delta(def.Style, style.Invalid, "input/input", def.CustomStyle)
}

func (def D) message() message.D {
	m := *def.Message
	m.InputName = def.Name
	m.Size = def.Size
	m.Style = def.style()
	m.CustomStyle = def.CustomStyle
	return m
}
