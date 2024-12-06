package cell

import "github.com/joaorufino/templui/components/style"

var DEFAULTS = D{
	Header:  true,
	Content: []string{"one", "two", "three"},
	ColSpan: "",
	CustomStyle: style.Custom{
		"cell": style.D{style.Add("text-sm")},
	},
}

type D struct {
	Header  bool
	Content any
	ColSpan string
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"table/cell":     style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) class() string {
	return style.CSSClass(style.Default, "table/cell", def.CustomStyle)
}
