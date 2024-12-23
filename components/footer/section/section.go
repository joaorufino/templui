package section

import (
	"github.com/joaorufino/templui/components/a"
	"github.com/joaorufino/templui/components/style"
)

var DEFAULTS = D{}

func init() {
	style.SetDefaults(style.Defaults{
		"footer/section/title": {
			style.Default: {
				style.Set("mb-6 text-sm font-semibold uppercase"),
			},
		},
		"footer/section/link": {
			style.Default: {
				style.Set("font-medium mb-4"),
			},
		},
	})
}

type D struct {
	Title string
	Links []a.D
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"footer/section/link":  style.D{style.Add("...")},
	// 		"footer/section/title": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) class(k string) string {
	return style.CSSClass(style.Default, k, def.CustomStyle)
}
