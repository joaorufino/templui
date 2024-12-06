package footer

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/social"

	"github.com/joaorufino/templui/components/style"
)

func AddfooterStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("footer",
		func(
			styleValue style.Style,
			copyright templ.Component,
			brand templ.Component,
			sections any,
			social []social.D,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Style: styleValue,

				Copyright: copyright,

				Brand: brand,

				Sections: sections,

				Social: social,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.ObjectArg("copyright", defaults.Copyright),

		storybook.ObjectArg("brand", defaults.Brand),

		storybook.ObjectArg("sections", defaults.Sections),

		storybook.SliceArg("social", defaults.Social),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
