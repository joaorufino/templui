package element

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/style"
)

func AddelementStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("element",
		func(
			id string,
			open bool,
			title string,
			content templ.Component,
			customstyle style.Custom,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				ID: id,

				Open: open,

				Title: title,

				Content: content,

				CustomStyle: customstyle,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.TextArg("id", defaults.ID),

		storybook.BooleanArg("open", defaults.Open),

		storybook.TextArg("title", defaults.Title),

		storybook.ObjectArg("content", defaults.Content),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
