package section

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/joaorufino/templui/components/a"

	"github.com/joaorufino/templui/components/style"
)

func AddsectionStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("section",
		func(
			title string,
			links []a.D,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				Title: title,

				Links: links,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.TextArg("title", defaults.Title),

		storybook.SliceArg("links", defaults.Links),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
