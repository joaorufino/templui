package checkboxgroup

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/joaorufino/templui/components/checkbox"

	"github.com/joaorufino/templui/components/form/validation/message"

	"github.com/joaorufino/templui/components/style"
)

func AddcheckboxgroupStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("checkboxgroup",
		func(
			name string,
			styleValue style.Style,
			checkboxes []checkbox.D,
			message message.D,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Name: name,

				Style: styleValue,

				Checkboxes: checkboxes,

				Message: &message,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.TextArg("name", defaults.Name),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.SliceArg("checkboxes", defaults.Checkboxes),

		storybook.ComponentArg("message", message.DEFAULTS),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
