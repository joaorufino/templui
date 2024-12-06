package control

import (
	"github.com/joaorufino/templui/components/icon"
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"
)

func AddcontrolStory(s *storybook.Storybook) {

	s.AddComponent("control",
		func(
			sidebarid string,
			icon any,
		) templ.Component {
			def := D{

				SidebarID: sidebarid,

				Icon: icon,
			}
			return C(def)
		},

		storybook.TextArg("sidebarid", ""),

		storybook.ObjectArg("icon", icon.C(icon.DEFAULTS)),
	)
}
