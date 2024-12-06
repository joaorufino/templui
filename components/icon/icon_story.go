package icon

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/size"

	"github.com/joaorufino/templui/components/style"
)

func AddiconStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("icon",
		func(
			icon string,
			styleValue style.Style,
			sizeValue size.Size,
			customstyle style.Custom,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				Icon: icon,

				Style: styleValue,

				Size: sizeValue,

				CustomStyle: customstyle,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.TextArg("icon", defaults.Icon),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.ObjectArg("sizeValue", defaults.Size),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
