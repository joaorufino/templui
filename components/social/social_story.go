package social

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/joaorufino/templui/components/size"

	"github.com/joaorufino/templui/components/style"
)

func AddsocialStory(s *storybook.Storybook) {

	s.AddComponent("social",
		func(
			typeValue Type,
			link string,
			sizeValue size.Size,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Type: typeValue,

				Link: link,

				Size: sizeValue,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.ObjectArg("typeValue", Type(0)),

		storybook.TextArg("link", ""),

		storybook.ObjectArg("sizeValue", size.Size(0)),

		storybook.CustomStyleArg("customstyle", style.Custom{}),
	)
}
