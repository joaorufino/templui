package checkbox

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/style"
)

func AddcheckboxStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("checkbox",
		func(
			id string,
			name string,
			label any,
			value string,
			checked bool,
			disabled bool,
			customstyle style.Custom,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				ID: id,

				Name: name,

				Label: label,

				Value: value,

				Checked: checked,

				Disabled: disabled,

				CustomStyle: customstyle,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.TextArg("id", defaults.ID),

		storybook.TextArg("name", defaults.Name),

		storybook.ObjectArg("label", defaults.Label),

		storybook.TextArg("value", defaults.Value),

		storybook.BooleanArg("checked", defaults.Checked),

		storybook.BooleanArg("disabled", defaults.Disabled),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
