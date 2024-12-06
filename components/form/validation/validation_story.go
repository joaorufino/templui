package validation

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"
)

func AddvalidationStory(s *storybook.Storybook) {

	s.AddComponent("validation",
		func(
			formid string,
			errors map[string]string,
		) templ.Component {
			def := D{

				FormID: formid,

				Errors: errors,
			}
			return C(def)
		},

		storybook.TextArg("formid", ""),

		storybook.ObjectArg("errors", map[string]string{}),
	)
}
