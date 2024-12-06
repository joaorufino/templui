package dropdown

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/a"

	"github.com/joaorufino/templui/components/button"

	"github.com/joaorufino/templui/components/style"
)

func AdddropdownStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("dropdown",
		func(
			button button.D,
			header templ.Component,
			links [][]a.D,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Button: button,

				Header: header,

				Links: links,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.ComponentArg("button", button.DEFAULTS),

		storybook.ObjectArg("header", defaults.Header),

		storybook.SliceArg("links", defaults.Links),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
