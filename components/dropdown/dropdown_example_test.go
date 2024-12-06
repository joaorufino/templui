package dropdown_test

import (
	"context"
	"os"

	"github.com/a-h/templ"
	"github.com/joaorufino/templui/components/a"
	"github.com/joaorufino/templui/components/button"
	"github.com/joaorufino/templui/components/dropdown"
)

func usernameComponent() templ.Component {
	return templ.Raw("User Name")
}

func ExampleC() {
	c := dropdown.C(dropdown.D{
		Button: button.D{Label: "Menu"},
		Header: usernameComponent(),
		Links: [][]a.D{{
			{Href: "/profile", Text: "Profile"},
		}, {
			{Href: "/logout", Text: "Logout"},
		}},
	})

	_ = c.Render(context.TODO(), os.Stdout)
}
