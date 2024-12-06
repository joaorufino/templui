package div

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/style"
)

func AdddivStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("div",
		func(
			content templ.Component,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Content: content,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.ObjectArg("content", defaults.Content),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
