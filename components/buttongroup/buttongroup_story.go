package buttongroup

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/button"

	"github.com/joaorufino/templui/components/size"

	"github.com/joaorufino/templui/components/style"
)

func AddbuttongroupStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("buttongroup",
		func(
			buttons []button.D,
			sizeValue size.Size,
			customstyle style.Custom,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				Buttons: buttons,

				Size: sizeValue,

				CustomStyle: customstyle,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.SliceArg("buttons", defaults.Buttons),

		storybook.ObjectArg("sizeValue", defaults.Size),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
