package radiogroup

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/joaorufino/templui/components/form/validation/message"

	"github.com/joaorufino/templui/components/radio"

	"github.com/joaorufino/templui/components/style"
)

func AddradiogroupStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("radiogroup",
		func(
			name string,
			styleValue style.Style,
			radios []radio.D,
			message message.D,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Name: name,

				Style: styleValue,

				Radios: radios,

				Message: &message,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.TextArg("name", defaults.Name),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.SliceArg("radios", defaults.Radios),

		storybook.ComponentArg("message", message.DEFAULTS),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
