package main

import (
	"log"
	"os"
	"text/template"

	"github.com/heaths/templ/formatter"
)

type Book struct {
	Title  string
	ISBN13 string
}

type Author struct {
	Name  string
	Books []Book
}

func main() {
	homes := []Author{
		{
			Name: "Stephen King",
			Books: []Book{
				{Title: "Misery: A Novel", ISBN13: "978-1501143106"},
				{Title: "Salem's Lot", ISBN13: "978-0345806796"},
				{Title: "The Dark Tower I: The Gunslinger", ISBN13: "978-1501143519"},
				{Title: "The Dark Tower II: The Drawing of the Three", ISBN13: "978-1501143533"},
			},
		},
		{
			Name: "Douglas Adams",
			Books: []Book{
				{Title: "The Hitchhiker's Guide to the Galaxy", ISBN13: "978-0345418913"},
				{Title: "The Restaurant at the End of the Universe", ISBN13: "978-0345418920"},
				{Title: "Life, the Universe and Everything", ISBN13: "978-0345418906"},
			},
		},
	}

	c := formatter.SystemColors()
	funcs := template.FuncMap{
		"green":  c.Green,
		"red":    c.Red,
		"yellow": c.Yellow,

		"truncate": formatter.Truncate,
	}

	t, err := template.New("main").Funcs(funcs).Parse(`{{range .}}{{with .Books}}{{template "table" .}}{{end}}{{end}}`)
	if err != nil {
		log.Fatalf("failed to compile %s template: %s", "table", err)
	}

	t, err = t.New("table").Parse(`{{range .}}{{truncate .Title 32 | printf "%-32s" | green}} {{.ISBN13}}{{"\n"}}{{end}}`)
	if err != nil {
		log.Fatalf("failed to compile %s template: %s", "main", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "main", homes)
	if err != nil {
		log.Fatalf("failed to execute template: %s", err)
	}
}
