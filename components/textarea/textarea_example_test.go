package textarea_test

import (
	"context"
	"os"

	"github.com/joaorufino/templui/components/icon"
	"github.com/joaorufino/templui/components/size"
	"github.com/joaorufino/templui/components/textarea"
)

func ExampleC() {
	c := textarea.C(textarea.D{
		Name:  "comment",
		Label: "Commentaire",
		Value: "Previous comment",
		Rows:  3,
		Size:  size.S,
		Icon:  icon.Text,
	})
	_ = c.Render(context.TODO(), os.Stdout)
}
