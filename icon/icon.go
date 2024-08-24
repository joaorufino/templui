//go:generate go run .build/generate.go
package icon

import "github.com/jfbus/templ-components/size"

// IconDefinition is the definition of icons.
// Usage: @icon.Icon(icon.IconDefinition{Icon:icon.Flower}) or @icon.Icon(icon.Flower) or &icon.Icon("flower")
type IconDefinition struct {
	// Icon is the SVG string (e.g. icon.Flower).
	Icon string
	// Class is an option CSS class to apply to the SVG tag.
	Class string
	// Size is the icon size.
	Size size.Size
}
