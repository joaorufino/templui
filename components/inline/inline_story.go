package inline

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/size"
)

func AddinlineStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("inline",
		func(
			value string,
			iconsize size.Size,
			edit templ.Component,
			defaultedit bool,
			focus string,
		) templ.Component {
			def := D{

				Value: value,

				IconSize: iconsize,

				Edit: edit,

				DefaultEdit: defaultedit,

				Focus: focus,
			}
			return C(def)
		},

		storybook.TextArg("value", defaults.Value),

		storybook.ObjectArg("iconsize", defaults.IconSize),

		storybook.ObjectArg("edit", defaults.Edit),

		storybook.BooleanArg("defaultedit", defaults.DefaultEdit),

		storybook.TextArg("focus", defaults.Focus),
	)
}
