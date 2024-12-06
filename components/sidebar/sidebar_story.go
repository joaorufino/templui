package sidebar

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/style"
)

func AddsidebarStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("sidebar",
		func(
			id string,
			content templ.Component,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				ID: id,

				Content: content,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.TextArg("id", defaults.ID),

		storybook.ObjectArg("content", defaults.Content),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
