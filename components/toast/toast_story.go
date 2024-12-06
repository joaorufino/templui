package toast

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/internal/storybook"

	"time"

	"github.com/joaorufino/templui/components/style"
)

func AddtoastStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("toast",
		func(
			_id string,
			containerid string,
			styleValue style.Style,
			icon string,
			content string,
			close Close,
			autoclosedelay time.Duration,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				_id: _id,

				ContainerID: containerid,

				Style: styleValue,

				Icon: icon,

				Content: content,

				Close: close,

				AutoCloseDelay: autoclosedelay,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.TextArg("_id", defaults._id),

		storybook.TextArg("containerid", defaults.ContainerID),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.TextArg("icon", defaults.Icon),

		storybook.TextArg("content", defaults.Content),

		storybook.ObjectArg("close", defaults.Close),

		storybook.ObjectArg("autoclosedelay", defaults.AutoCloseDelay),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
