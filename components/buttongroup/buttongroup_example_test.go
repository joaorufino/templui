package buttongroup_test

import (
	"context"
	"os"

	"github.com/a-h/templ"
	"github.com/joaorufino/templui/components/button"
	"github.com/joaorufino/templui/components/buttongroup"
	"github.com/joaorufino/templui/components/icon"
	"github.com/joaorufino/templui/components/size"
	"github.com/joaorufino/templui/components/style"
)

func ExampleC() {
	c := buttongroup.C(buttongroup.D{
		Size: size.S,
		Buttons: []button.D{
			{
				Icon:  icon.ArrowDownNarrowWide,
				Label: "Sort",
				Style: button.StyleHideLabelAlways,
				CustomStyle: style.Custom{
					"button": style.D{
						style.ReplaceVariants("bg", "bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-600"),
					},
				},
				Attributes: templ.Attributes{
					"@click": "$dispatch('input', '')",
				},
			},
			{
				Icon:  icon.Heart,
				Label: "Rating",
				Style: button.StyleHideLabelAlways,
				Attributes: templ.Attributes{
					"@click": "$dispatch('input', 'rating')",
				},
			},
			{
				Icon:  icon.Banknote,
				Label: "Price",
				Style: button.StyleHideLabelAlways,
				Attributes: templ.Attributes{
					"@click": "$dispatch('input', 'price')",
				},
			},
		},
		Attributes: templ.Attributes{
			"x-model": "sort",
		},
	})
	_ = c.Render(context.TODO(), os.Stdout)
}
