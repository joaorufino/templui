package table

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/joaorufino/templui/components/style"

	"github.com/joaorufino/templui/components/table/row"
)

func AddtableStory(s *storybook.Storybook) {
	s.AddComponent("table",
		func(
			styleValue style.Style,
			header *row.D,
			rows []row.D,
			footer *row.D,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Style: styleValue,

				Header: header,

				Rows: rows,

				Footer: footer,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.StyleArg("styleValue", DEFAULTS.Style),

		storybook.ComponentArg("header", DEFAULTS.Header),

		storybook.SliceArg("rows", DEFAULTS.Rows),

		storybook.ComponentArg("footer", DEFAULTS.Footer),

		storybook.CustomStyleArg("customstyle", DEFAULTS.CustomStyle),
	)
}
