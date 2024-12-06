package control

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/components/button"
	"github.com/joaorufino/templui/components/icon"
	"github.com/joaorufino/templui/components/sidebar"
	"github.com/joaorufino/templui/components/size"
	"github.com/joaorufino/templui/components/style"
)

type D struct {
	SidebarID string
	Icon      any
}

func (def D) sidebarId() string {
	if def.SidebarID != "" {
		return def.SidebarID
	}
	return sidebar.DefaultID
}

func (def D) button() button.D {
	return button.D{
		Icon:        icon.Menu,
		Style:       button.StyleOutline,
		CustomStyle: style.Custom{"button": {style.Set("md:hidden")}},
		Size:        size.L,
		Attributes: templ.Attributes{
			"@click.stop": "sidebar = !sidebar",
		},
	}
}
