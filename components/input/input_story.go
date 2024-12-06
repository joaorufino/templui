package input

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/form/validation/message"

	"github.com/joaorufino/templui/components/position"

	"github.com/joaorufino/templui/components/size"

	"github.com/joaorufino/templui/components/style"

	"github.com/joaorufino/templui/components/tooltip"
)

func AddinputStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("input",
		func(
			id string,
			name string,
			typeValue Type,
			styleValue style.Style,
			label any,
			value string,
			placeholder string,
			message message.D,
			disabled bool,
			invalid bool,
			sizeValue size.Size,
			loader bool,
			icon string,
			iconposition position.Position,
			tooltip tooltip.D,
			customstyle style.Custom,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				ID: id,

				Name: name,

				Type: typeValue,

				Style: styleValue,

				Label: label,

				Value: value,

				Placeholder: placeholder,

				Message: &message,

				Disabled: disabled,

				Invalid: invalid,

				Size: sizeValue,

				Loader: loader,

				Icon: icon,

				IconPosition: iconposition,

				Tooltip: &tooltip,

				CustomStyle: customstyle,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.TextArg("id", defaults.ID),

		storybook.TextArg("name", defaults.Name),

		storybook.ObjectArg("typeValue", defaults.Type),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.ObjectArg("label", defaults.Label),

		storybook.TextArg("value", defaults.Value),

		storybook.TextArg("placeholder", defaults.Placeholder),

		storybook.ComponentArg("message", message.DEFAULTS),

		storybook.BooleanArg("disabled", defaults.Disabled),

		storybook.BooleanArg("invalid", defaults.Invalid),

		storybook.ObjectArg("sizeValue", defaults.Size),

		storybook.BooleanArg("loader", defaults.Loader),

		storybook.TextArg("icon", defaults.Icon),

		storybook.ObjectArg("iconposition", defaults.IconPosition),

		storybook.ComponentArg("tooltip", tooltip.DEFAULTS),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
