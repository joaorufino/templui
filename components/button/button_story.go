package button

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/position"

	"github.com/joaorufino/templui/components/size"

	"github.com/joaorufino/templui/components/style"

	"github.com/joaorufino/templui/components/tooltip"
)

func AddbuttonStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("button",
		func(
			id string,
			typeValue Type,
			label string,
			styleValue style.Style,
			sizeValue size.Size,
			icon string,
			iconposition position.Position,
			disabled bool,
			loader bool,
			tooltip tooltip.D,
			stylekey string,
			customstyle style.Custom,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				ID: id,

				Type: typeValue,

				Label: label,

				Style: styleValue,

				Size: sizeValue,

				Icon: icon,

				IconPosition: iconposition,

				Disabled: disabled,

				Loader: loader,

				Tooltip: &tooltip,

				StyleKey: stylekey,

				CustomStyle: customstyle,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.TextArg("id", defaults.ID),

		storybook.ObjectArg("typeValue", defaults.Type),

		storybook.TextArg("label", defaults.Label),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.ObjectArg("sizeValue", defaults.Size),

		storybook.TextArg("icon", defaults.Icon),

		storybook.ObjectArg("iconposition", defaults.IconPosition),

		storybook.BooleanArg("disabled", defaults.Disabled),

		storybook.BooleanArg("loader", defaults.Loader),

		storybook.ComponentArg("tooltip", tooltip.DEFAULTS),

		storybook.TextArg("stylekey", defaults.StyleKey),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
