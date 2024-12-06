package dropzone

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/form/validation/message"

	"github.com/joaorufino/templui/components/icon"

	"github.com/joaorufino/templui/components/size"

	"github.com/joaorufino/templui/components/style"
)

func AdddropzoneStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("dropzone",
		func(
			id string,
			name string,
			multiple bool,
			allowedtypes []string,
			styleValue style.Style,
			icon icon.D,
			dropicon icon.D,
			iconsize size.Size,
			dragmessage string,
			dropmessage string,
			allowedtypesmessage string,
			message message.D,
			disabled bool,
			invalid bool,
			loader bool,
			customstyle style.Custom,
			attributes templ.Attributes,
		) templ.Component {
			def := D{

				ID: id,

				Name: name,

				Multiple: multiple,

				AllowedTypes: allowedtypes,

				Style: styleValue,

				Icon: &icon,

				DropIcon: &dropicon,

				IconSize: iconsize,

				DragMessage: dragmessage,

				DropMessage: dropmessage,

				AllowedTypesMessage: allowedtypesmessage,

				Message: &message,

				Disabled: disabled,

				Invalid: invalid,

				Loader: loader,

				CustomStyle: customstyle,

				Attributes: attributes,
			}
			return C(def)
		},

		storybook.TextArg("id", defaults.ID),

		storybook.TextArg("name", defaults.Name),

		storybook.BooleanArg("multiple", defaults.Multiple),

		storybook.TextArg("allowedtypes", ""),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.ComponentArg("icon", icon.DEFAULTS),

		storybook.ComponentArg("dropicon", icon.DEFAULTS),

		storybook.ObjectArg("iconsize", defaults.IconSize),

		storybook.TextArg("dragmessage", defaults.DragMessage),

		storybook.TextArg("dropmessage", defaults.DropMessage),

		storybook.TextArg("allowedtypesmessage", defaults.AllowedTypesMessage),

		storybook.ComponentArg("message", message.DEFAULTS),

		storybook.BooleanArg("disabled", defaults.Disabled),

		storybook.BooleanArg("invalid", defaults.Invalid),

		storybook.BooleanArg("loader", defaults.Loader),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),

		storybook.ObjectArg("attributes", defaults.Attributes),
	)
}
