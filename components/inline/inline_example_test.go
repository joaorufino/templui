package inline_test

import (
	"context"
	"os"

	"github.com/a-h/templ"
	"github.com/joaorufino/templui/components/icon"
	"github.com/joaorufino/templui/components/inline"
	"github.com/joaorufino/templui/components/input"
	"github.com/joaorufino/templui/components/label"
	"github.com/joaorufino/templui/components/position"
	"github.com/joaorufino/templui/components/size"
)

func ExampleC() {
	c := inline.C(inline.D{
		Value:    "Previous value",
		IconSize: size.S,
		Focus:    "itemtitle",
		Edit: input.C(input.D{
			Name:         "title",
			Label:        label.D{Label: "Title", Hide: true},
			Value:        "Previous value",
			Icon:         icon.CornerDownLeft,
			IconPosition: position.End,
			Size:         size.S,
			Loader:       true,
			Attributes: templ.Attributes{
				"hx-post":   "/items/update_title/24",
				"hx-target": "#item_24",
				"x-ref":     "itemtitle",
			},
		}),
	})
	c.Render(context.TODO(), os.Stdout)
}
