package body

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/footer"

	"github.com/joaorufino/templui/components/navbar"

	"github.com/joaorufino/templui/components/sidebar"

	"github.com/joaorufino/templui/components/style"

	"github.com/joaorufino/templui/components/toast/container"
)

func AddbodyStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("body",
		func(
			navbar navbar.D,
			sidebar sidebar.D,
			toasts container.D,
			footer footer.D,
			state map[string]string,
			navbarheight NavbarHeight,
			customstyle style.Custom,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				Navbar: &navbar,

				Sidebar: &sidebar,

				Toasts: &toasts,

				Footer: &footer,

				State: state,

				NavbarHeight: navbarheight,

				CustomStyle: customstyle,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.ComponentArg("navbar", navbar.DEFAULTS),

		storybook.ComponentArg("sidebar", sidebar.DEFAULTS),

		storybook.ComponentArg("toasts", container.DEFAULTS),

		storybook.ComponentArg("footer", footer.DEFAULTS),

		storybook.ObjectArg("state", defaults.State),

		storybook.ObjectArg("navbarheight", defaults.NavbarHeight),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
