// Package radio implements radio fields.
package radio

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/helper"
	"github.com/jfbus/templ-components/components/label"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

const (
	StyleBordered        style.Style = 2
	StyleGrouped         style.Style = 4
	StyleGroupedVertical style.Style = 8
	StyleLabelOnly       style.Style = 16
)

// Defaults defines the default Color/Class. They are overriden by D.Class.
var Defaults = style.Defaults{
	style.StyleDefault: {
		"ContainerClass": style.D{},
		"InputClass": style.D{
			Class: "w-4 h-4 focus:ring-2",
			Color: "text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:bg-gray-700 dark:border-gray-600",
		},
	},
	StyleBordered: {
		"ContainerClass": {
			Color: "border-gray-200 dark:border-gray-700",
			Class: "flex items-center border rounded",
		},
	},
	StyleGrouped: {
		"ContainerClass": {
			Class: "border-b rounded-t-lg last:border-0 flex items-center ps-3",
		},
	},
	StyleGroupedVertical: {
		"ContainerClass": {
			Class: "border-b rounded-t-lg last:border-0 items-center ps-3",
		},
	},
	StyleLabelOnly: {
		"ContainerClass": {
			Class: "inline-flex items-center justify-between",
		},
		"InputClass": {
			Color: " ",
			Class: "hidden peer",
		},
		"LabelClass": {
			Class: "border p-2 rounded-lg cursor-pointer",
			Color: "text-gray-500 bg-white border-gray-200 dark:hover:text-gray-300 dark:border-gray-700 dark:peer-checked:text-blue-500 peer-checked:border-blue-600 peer-checked:text-blue-600 hover:text-gray-600 hover:bg-gray-100 dark:text-gray-400 dark:bg-gray-800 dark:hover:bg-gray-700",
		},
	},
}

// D is the definition for input fields.
type D struct {
	// ID is the input id (Name if not set).
	ID string
	// Name is the input name.
	Name string
	// Style is the input style.
	Style style.Style
	// Label is the input label.
	Label any
	// Value is the input value.
	Value   string
	Checked bool
	// Disabled disables the input.
	Disabled bool
	// Size defines the input size (size.S, size.Normal (default) or size.L).
	Size size.Size
	// Container
	ContainerClass style.D
	InputClass     style.D
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) label() label.D {
	defaults := style.Default(Defaults, def.Style, "LabelClass")
	switch l := def.Label.(type) {
	case string:
		return label.D{
			InputID: def.id(),
			Label:   l,
			Class:   defaults,
		}
	case templ.Component:
		return label.D{
			InputID: def.id(),
			Label:   l,
			Class:   defaults,
		}
	case label.D:
		l.InputID = def.id()
		l.Class.Color = helper.IfEmpty(l.Class.Color, defaults.Color)
		l.Class.Class = helper.IfEmpty(l.Class.Class, defaults.Class)
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

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(style.Default(Defaults, def.Style, "ContainerClass"))
}

func (def D) inputClass() string {
	return def.InputClass.CSSClass(style.Default(Defaults, def.Style, "InputClass"))
}
