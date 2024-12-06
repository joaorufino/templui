package div

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/components/style"
)

var DEFAULTS = D{}

type D struct {
	//playground:import:github.com/joaorufino/templui/components/button
	//playground:default:button.C(button.D{Label:"Button"})
	Content     templ.Component
	CustomStyle style.Custom
}

func (def D) class() string {
	return style.CSSClass(style.Default, "div", def.CustomStyle)
}
