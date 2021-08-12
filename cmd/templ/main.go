package main

import (
	"log"
	"os"

	"github.com/heaths/templ/internal/catalog"
)

func main() {
	authors := catalog.New()
	if err := catalog.Format(authors, os.Stdout); err != nil {
		log.Fatalf("failed to format catalog: %s", err)
	}
}
