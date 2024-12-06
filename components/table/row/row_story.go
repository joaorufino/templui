package row

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/style"
)

func AddrowStory(s *storybook.Storybook) {

	s.AddComponent("row",
		func(
			header bool,
			cells any,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Header: header,

				Cells: cells,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.BooleanArg("header", DEFAULTS.Header),

		storybook.ObjectArg("cells", DEFAULTS.Cells),

		storybook.CustomStyleArg("customstyle", DEFAULTS.CustomStyle),
	)
}
