package navbar

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/style"
)

func AddnavbarStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("navbar",
		func(
			styleValue style.Style,
			sections []templ.Component,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Style: styleValue,

				Sections: sections,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.SliceArg("sections", defaults.Sections),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
