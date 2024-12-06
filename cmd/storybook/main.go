package main

import (
	"context"
	"log"

	"github.com/joaorufino/templui/components/table"

	"github.com/joaorufino/templui/components/a"
	"github.com/joaorufino/templui/internal/storybook"
)

func Stories(s *storybook.Storybook) {
	a.AddaStory(s)
	table.AddtableStory(s)

}

func main() {
	s := storybook.New()
	s.WithAdditionalPreviewJS("import '../static/css/tailwind.css';import '../static/script/alpine.js';")
	Stories(s)

	ctx := context.Background()
	if err := s.ListenAndServeWithContext(ctx); err != nil {
		log.Fatal(err)
	}
}
