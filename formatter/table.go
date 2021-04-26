package formatter

import (
	"io"
	"strings"
	"text/tabwriter"
)

type Table struct {
	w *tabwriter.Writer
}

func NewTable(output io.Writer) Table {
	return Table{
		w: tabwriter.NewWriter(output, 10, 1, 1, ' ', 0),
	}
}

func (t *Table) TableRow(fields ...string) string {
	row := strings.Join(fields, "\t")
	t.w.Write([]byte(row + "\n"))

	return ""
}

func (t *Table) EndTable() string {
	t.w.Flush()

	return ""
}
