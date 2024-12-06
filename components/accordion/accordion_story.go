package accordion

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/joaorufino/templui/components/accordion/element"

	"github.com/joaorufino/templui/components/style"
)

func AddaccordionStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("accordion",
		func(
			id string,
			children []element.D,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				ID: id,

				Children: children,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.TextArg("id", defaults.ID),

		storybook.SliceArg("children", defaults.Children),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
