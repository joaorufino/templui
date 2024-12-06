package cell

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/joaorufino/templui/components/style"
)

func AddcellStory(s *storybook.Storybook) {

	s.AddComponent("cell",
		func(
			header bool,
			content any,
			colspan string,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Header: header,

				Content: content,

				ColSpan: colspan,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.BooleanArg("header", DEFAULTS.Header),

		storybook.ObjectArg("content", DEFAULTS.Content),

		storybook.TextArg("colspan", DEFAULTS.ColSpan),

		storybook.CustomStyleArg("customstyle", DEFAULTS.CustomStyle),
	)
}
