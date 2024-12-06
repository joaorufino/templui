package container

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/joaorufino/templui/components/style"
)

func AddcontainerStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("container",
		func(
			id string,
			styleValue style.Style,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				ID: id,

				Style: styleValue,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.TextArg("id", defaults.ID),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
