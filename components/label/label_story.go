package label

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/style"

	"github.com/joaorufino/templui/components/tooltip"
)

func AddlabelStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("label",
		func(
			inputid string,
			inputname string,
			styleValue style.Style,
			label any,
			hide bool,
			novalidation bool,
			tooltip tooltip.D,
			customstyle style.Custom,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				InputID: inputid,

				InputName: inputname,

				Style: styleValue,

				Label: label,

				Hide: hide,

				NoValidation: novalidation,

				Tooltip: &tooltip,

				CustomStyle: customstyle,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.TextArg("inputid", defaults.InputID),

		storybook.TextArg("inputname", defaults.InputName),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.ObjectArg("label", defaults.Label),

		storybook.BooleanArg("hide", defaults.Hide),

		storybook.BooleanArg("novalidation", defaults.NoValidation),

		storybook.ComponentArg("tooltip", tooltip.DEFAULTS),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
