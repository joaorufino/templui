package form

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"
)

func AddformStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("form",
		func(
			id string,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				ID: id,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.TextArg("id", defaults.ID),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
