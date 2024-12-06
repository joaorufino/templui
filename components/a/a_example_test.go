package a_test

import (
	"context"
	"os"

	"github.com/joaorufino/templui/components/style"
	"github.com/joaorufino/templui/components/a"
)

func ExampleC() {
	c := a.C(a.D{
		Href: "https://www.example.com",
		Text: "Example",
		CustomStyle: style.Custom{
			"a": style.D{style.Add("text-sm")},
		},
	})
	_ = c.Render(context.TODO(), os.Stdout)
	// output: <a href="https://www.example.com" class="hover:underline text-sm">Example</a>
}
