package icon_test

import (
	"context"
	"os"

	"github.com/joaorufino/templui/components/icon"
)

func ExampleC() {
	c := icon.C(icon.Flower)
	_ = c.Render(context.TODO(), os.Stdout)
}
