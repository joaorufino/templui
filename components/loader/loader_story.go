package loader

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/joaorufino/templui/components/size"

	"github.com/joaorufino/templui/components/style"
)

func AddloaderStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("loader",
		func(
			sizeValue size.Size,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Size: sizeValue,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.ObjectArg("sizeValue", defaults.Size),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
