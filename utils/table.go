package utils

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

type Table struct {
	Headers []string
	Rows    [][]string
}

func (t *Table) String() string {
	var buffer bytes.Buffer
	w := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)
	for _, header := range t.Headers {
		fmt.Fprintf(w, "%s\t", header)
	}
	fmt.Fprintln(w)
	for _, row := range t.Rows {
		for _, cell := range row {
			fmt.Fprintf(w, "%s\t", cell)
		}
		fmt.Fprintln(w)
	}
	w.Flush()
	return buffer.String()
}
