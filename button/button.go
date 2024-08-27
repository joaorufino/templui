// Package button implements buttons.
package button

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/position"
	"github.com/jfbus/templ-components/size"
)

type Type string

const (
	Button Type = "button"
	Submit Type = "submit"
)

type Style int

const (
	Normal   Style = 0
	Outline  Style = 1
	Pill     Style = 2
	NoBorder Style = 4
)

const (
	DefaultOutlineColor = "text-gray-900 bg-white border-gray-300 hover:bg-gray-100 focus:ring-gray-100 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"
	DefaultColor        = "text-white bg-blue-700 hover:bg-blue-800 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
)

// D is the button definition.
type D struct {
	// ID is the button id.
	ID string
	// Type is the button type (button, submit, ...).
	Type Type
	// Label is the button label.
	Label string
	// HideLabel defines if the label is only available to screenreaders.
	HideLabel bool
	// Style is the button style (Normal, Pill, Outline ou Pill|Outline).
	Style Style
	// Color overrides the default color CSS classes (DefaultOutlineColor if Outline, DefaultColor otherwise)
	// e.g.
	// - for a plain green button: text-white bg-green-700 hover:bg-green-800 focus:ring-green-300 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800
	// - for a outline green button: text-green-700 hover:text-white border-green-700 hover:bg-green-800 focus:ring-green-300 dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800
	Color string
	// Size defines the input size (size.XS, size.S, size.Normal (default) size.L or size.XL).
	Size size.Size
	// Icon displays an icon on the left side.
	Icon string
	// IconPosition can be position.Start (default) or position.End.
	IconPosition position.Position
	// Disabled disables the input.
	Disabled bool
	// Loader displays a spinning loader when an HTMX action is triggered by the input.
	Loader bool
	// Class defines additional CSS class for the button.
	Class string
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) buttonType() string {
	if def.Type == "" {
		return string(Button)
	}
	return string(def.Type)
}

func (def D) buttonClass() string {
	class := "focus:ring-4 focus:outline-none"
	if def.Style&Outline != 0 {
		class += " border"
		if def.Color == "" {
			class += " " + DefaultOutlineColor
		}
	} else if def.Color == "" {
		class += " " + DefaultColor
	}

	if def.Style&Pill != 0 {
		class += " rounded-full"
	} else if def.Style&NoBorder == 0 {
		class += " rounded-lg"
	}
	switch {
	case def.noLabel() && def.Size >= size.Normal:
		class += " p-2.5 text-sm"
	case def.noLabel() && def.Size == size.XS:
		class += " p-1 text-sm"
	case def.noLabel():
		class += " p-1.5 text-sm"
	case def.Size == size.XS:
		class += " px-3 py-2 text-xs"
	case def.Size == size.S:
		class += " px-3 py-2 text-sm"
	case def.Size == size.L:
		class += " px-5 py-3 text-base"
	case def.Size == size.XL:
		class += " px-6 py-3.5 text-base"
	default:
		class += " px-5 py-2.5 text-base"
	}
	if def.Icon != "" || def.Loader {
		class += " inline-flex items-center"
	}
	if def.Color != "" {
		class += " " + def.Color
	}
	if def.Class != "" {
		class += " " + def.Class
	}
	return class
}

func (def D) noLabel() bool {
	return def.HideLabel || def.Label == ""
}

func (def D) iconSize() size.Size {
	if def.HideLabel || def.Label == "" {
		return def.Size + 1
	}
	return def.Size
}

func (def D) iconClass() string {
	if def.noLabel() {
		return ""
	}
	if def.IconPosition == position.End {
		return "ms-2"
	}
	return "me-2"
}
