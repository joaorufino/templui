package modal

import (
	"github.com/joaorufino/templui/internal/storybook"

	"github.com/a-h/templ"

	"github.com/joaorufino/templui/components/button"

	"github.com/joaorufino/templui/components/form"

	"github.com/joaorufino/templui/components/style"
)

func AddmodalStory(s *storybook.Storybook) {

	// Use component defaults where available
	defaults := DEFAULTS

	s.AddComponent("modal",
		func(
			id string,
			styleValue style.Style,
			maxwidth MaxWidth,
			title string,
			form form.D,
			content any,
			close button.D,
			confirm button.D,
			customstyle style.Custom,
		) templ.Component {
			def := D{

				ID: id,

				Style: styleValue,

				MaxWidth: maxwidth,

				Title: title,

				Form: &form,

				Content: content,

				Close: &close,

				Confirm: &confirm,

				CustomStyle: customstyle,
			}
			return C(def)
		},

		storybook.TextArg("id", defaults.ID),

		storybook.StyleArg("styleValue", defaults.Style),

		storybook.ObjectArg("maxwidth", defaults.MaxWidth),

		storybook.TextArg("title", defaults.Title),

		storybook.ComponentArg("form", form.DEFAULTS),

		storybook.ObjectArg("content", defaults.Content),

		storybook.ComponentArg("close", button.DEFAULTS),

		storybook.ComponentArg("confirm", button.DEFAULTS),

		storybook.CustomStyleArg("customstyle", defaults.CustomStyle),
	)
}
