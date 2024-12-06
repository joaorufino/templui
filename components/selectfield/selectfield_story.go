package selectfield

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/form/validation/message"

	"github.com/joaorufino/templui/components/selectfield/option"

	"github.com/joaorufino/templui/components/size"

	"github.com/joaorufino/templui/components/style"

	"github.com/joaorufino/templui/components/tooltip"
)

func AddselectfieldStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("selectfield",
		func(
			id string,
			name string,
			styleValue style.Style,
			label any,
			options []option.D,
			selected string,
			disabled bool,
			sizeValue size.Size,
			message message.D,
			tooltip tooltip.D,
			customstyle style.Custom,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				ID: id,

				Name: name,

				Style: styleValue,

				Label: label,

				Options: options,

				Selected: selected,

				Disabled: disabled,

				Size: sizeValue,

				Message: &message,

				Tooltip: &tooltip,

				CustomStyle: customstyle,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.TextArg("id", defaults.ID),

		storybook.TextArg("name", defaults.Name),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.ObjectArg("label", defaults.Label),

		storybook.SliceArg("options", defaults.Options),

		storybook.TextArg("selected", defaults.Selected),

		storybook.BooleanArg("disabled", defaults.Disabled),

		storybook.ObjectArg("sizeValue", defaults.Size),

		storybook.ComponentArg("message", message.DEFAULTS),

		storybook.ComponentArg("tooltip", tooltip.DEFAULTS),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
