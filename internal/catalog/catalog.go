package catalog

import (
	"fmt"
	"io"
	"text/template"

	"github.com/heaths/templ/pkg/formatter"
)

type Book struct {
	Title  string
	ISBN13 string
}

type Author struct {
	Name  string
	Books []Book
}

func New() []Author {
	return []Author{
		{
			Name: "Stephen King",
			Books: []Book{
				{Title: `Misery: A Novel`, ISBN13: "978-1501143106"},
				{Title: `Salem's Lot`, ISBN13: "978-0345806796"},
				{Title: `The Dark Tower I: The Gunslinger`, ISBN13: "978-1501143519"},
				{Title: `The Dark Tower II: The Drawing of the Three`, ISBN13: "978-1501143533"},
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
}

func Format(authors []Author, output io.Writer) error {
	c := formatter.SystemColors()
	table := formatter.NewTable(output)

	funcs := template.FuncMap{
		"green":  c.Green,
		"red":    c.Red,
		"yellow": c.Yellow,

		"tableRow": table.TableRow,
		"endTable": table.EndTable,

		"truncate": formatter.Truncate,
	}

	t, err := template.New("main").Funcs(funcs).Parse(`{{range .}}{{with .Books}}{{template "table" .}}{{end}}{{end}}{{endTable}}`)
	if err != nil {
		return fmt.Errorf("failed to compile %s template: %s", "table", err)
	}

	t, err = t.New("table").Parse(`{{range .}}{{tableRow (truncate .Title 30 | green) (.ISBN13 | yellow)}}{{end}}`)
	if err != nil {
		return fmt.Errorf("failed to compile %s template: %s", "main", err)
	}

	err = t.ExecuteTemplate(output, "main", authors)
	if err != nil {
		return fmt.Errorf("failed to execute template: %s", err)
	}

	return nil
}
