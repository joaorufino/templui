package message

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/joaorufino/templui/components/size"

	"github.com/joaorufino/templui/components/style"
)

func AddmessageStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("message",
		func(
			inputname string,
			styleValue style.Style,
			message string,
			sizeValue size.Size,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				InputName: inputname,

				Style: styleValue,

				Message: message,

				Size: sizeValue,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.TextArg("inputname", defaults.InputName),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.TextArg("message", defaults.Message),

		storybook.ObjectArg("sizeValue", defaults.Size),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
