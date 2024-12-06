package input_test

import (
	"context"
	"os"

	"github.com/joaorufino/templui/components/icon"
	"github.com/joaorufino/templui/components/input"
	"github.com/joaorufino/templui/components/size"
)

func ExampleC() {
	c := input.C(input.D{
		Name:        "title",
		Type:        input.TypeText,
		Label:       "Title",
		Value:       "Previous title",
		Placeholder: "Enter a title",
		Size:        size.S,
		Icon:        icon.Bookmark,
	})
	_ = c.Render(context.TODO(), os.Stdout)
}
