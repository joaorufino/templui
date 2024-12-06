package breadcrumb

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/components/style"
)

// Initialize default styles
var DEFAULTS = D{}

func init() {
	style.SetDefaults(
		style.Defaults{
			"breadcrumb": {
				style.Default: {
					style.Set("flex items-center text-sm font-medium"),
				},
				StyleDark: {
					style.Add("text-gray-400"),
				},
				style.Default | StyleContainer: {
					style.Set("flex items-center"),
				},
				style.Default | StyleItem: {
					style.Set("inline-flex items-center"),
				},
				style.Default | StyleSeparator: {
					style.Set("mx-2 text-gray-400"),
				},
				style.Default | StyleLink: {
					style.Set("text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-white"),
				},
				style.Default | StyleCurrent: {
					style.Set("text-gray-700 dark:text-white"),
				},
			},
		},
	)
}

// Style constants
const (
	StyleDark      style.Style = 1 << 8
	StyleContainer style.Style = 1 << 9
	StyleItem      style.Style = 1 << 10
	StyleSeparator style.Style = 1 << 11
	StyleLink      style.Style = 1 << 12
	StyleCurrent   style.Style = 1 << 13
)

// Link represents a breadcrumb item
type Link struct {
	Text string
	Href templ.SafeURL
	Icon string
}

// Configuration struct
type D struct {
	Links  []Link
	Style  style.Style
	Custom style.Custom
}

// Style helper methods
func (def D) style() style.Style {
	if def.Style > 0 {
		return def.Style
	}
	return style.Default
}

func (def D) containerClass() string {
	return style.CSSClass(def.style()|StyleContainer, "breadcrumb", def.Custom)
}

func (def D) itemClass() string {
	return style.CSSClass(def.style()|StyleItem, "breadcrumb", def.Custom)
}

func (def D) separatorClass() string {
	return style.CSSClass(def.style()|StyleSeparator, "breadcrumb", def.Custom)
}

func (def D) linkClass() string {
	return style.CSSClass(def.style()|StyleLink, "breadcrumb", def.Custom)
}

func (def D) currentClass() string {
	return style.CSSClass(def.style()|StyleCurrent, "breadcrumb", def.Custom)
}
