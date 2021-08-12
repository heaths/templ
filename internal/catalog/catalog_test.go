package catalog

import (
	"bytes"
	"testing"
)

func TestFormat(t *testing.T) {
	authors := []Author{
		{
			Name: "Name",
			Books: []Book{
				{Title: "Title 1", ISBN13: "0123"},
				{Title: "Longer Title 2", ISBN13: "4567"},
			},
		},
	}

	want := "Title 1        0123\n" +
		"Longer Title 2 4567\n"

	t.Run("format catalog", func(t *testing.T) {
		buf := &bytes.Buffer{}
		if err := Format(authors, buf); err != nil {
			t.Errorf("failed to format catalog: %s", err)
		}
		if got := buf.String(); got != want {
			t.Errorf("want: %s, got: %s", want, got)
		}
	})
}
