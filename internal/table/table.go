package table

import (
	"fmt"
	"strings"
)

type Table struct {
	headers []string
	rows    [][]interface{}
}

func NewTable(headers ...string) *Table {
	return &Table{headers, nil}
}

func (t *Table) AddRow(values ...interface{}) {
	if len(values) != len(t.headers) {
		fmt.Println("Incorrect number of values for the table row.")
		return
	}
	t.rows = append(t.rows, values)
}

func (t *Table) String() string {
	var tableStr strings.Builder

	colWidths := make([]int, len(t.headers))
	for i, header := range t.headers {
		colWidths[i] = len(header)
	}
	for _, row := range t.rows {
		for i, value := range row {
			valStr := fmt.Sprintf("%v", value)
			if len(valStr) > colWidths[i] {
				colWidths[i] = len(valStr)
			}
		}
	}

	for i, header := range t.headers {
		tableStr.WriteString(fmt.Sprintf("%-*s", colWidths[i], header))
		if i < len(t.headers)-1 {
			tableStr.WriteString("  ")
		}
	}
	tableStr.WriteString("\n")

	for _, row := range t.rows {
		for i, value := range row {
			valStr := fmt.Sprintf("%v", value)
			tableStr.WriteString(fmt.Sprintf("%-*s", colWidths[i], valStr))
			if i < len(row)-1 {
				tableStr.WriteString("  ")
			}
		}
		tableStr.WriteString("\n")
	}

	return tableStr.String()
}
