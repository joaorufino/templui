package accordion_test

import (
	"github.com/a-h/templ"
	"github.com/joaorufino/templui/components/accordion"
	"github.com/joaorufino/templui/components/accordion/element"
)

var yourcomponent = func() templ.Component {
	return templ.Raw("Content")
}

func ExampleC() {
	accordion.C(accordion.D{
		ID: "accordion",
		Children: []element.D{{
			Open:    true,
			Title:   "First",
			Content: yourcomponent(),
		}, {
			Title:   "Second",
			Content: yourcomponent(),
		}},
	})
}
