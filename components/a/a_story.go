package a

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/style"
)

func AddaStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("a",
		func(
			href string,
			text string,
			customstyle style.Custom,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				Href: href,

				Text: text,

				CustomStyle: customstyle,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.TextArg("href", DEFAULTS.Href),

		storybook.TextArg("text", DEFAULTS.Text),

		storybook.CustomStyleArg("customstyle", DEFAULTS.CustomStyle),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
