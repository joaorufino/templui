package div_test

import (
	"context"
	"os"

	"github.com/a-h/templ"
	"github.com/joaorufino/templui/components/div"
	"github.com/joaorufino/templui/components/style"
)

func myComponent() templ.Component {
	return templ.Raw(`Content`)
}

func ExampleC() {
	c := div.C(div.D{
		Content: myComponent(),
		CustomStyle: style.Custom{
			"div": style.D{
				style.Set("text-sm"),
			},
		},
	})
	_ = c.Render(context.TODO(), os.Stdout)
	// output: <div class="text-sm">Content</div>
}
