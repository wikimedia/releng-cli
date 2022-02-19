package output

/*
Creates output list this:

```
Repository                           File
Extension:FileImporter               tests/phpunit/Data/SourceUrlTest.php
Extension:Wikibase                   repo/tests/phpunit/includes/Api/EditEntityTest.php
SecurityCheckPlugin                  tests/integration/redos/test.php
SecurityCheckPlugin                  tests/integration/redos/test.php
```
*/
import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Table struct {
	Headings []interface{}
	Rows     [][]interface{}
}

func NewTable(headings []interface{}, rows [][]interface{}) *Table {
	return &Table{
		Headings: headings,
		Rows:     rows,
	}
}

func TableFromObjects(objects []interface{}, headings []string, objectToRow func(interface{}) []string) *Table {
	var thisTable Table
	thisTable.AddHeadingsS(headings...)
	for _, object := range objects {
		thisTable.AddRow(stringSplitToInterfaceSplit(objectToRow(object))...)
	}
	return &thisTable
}

func (t *Table) AddHeadings(headings ...interface{}) {
	t.Headings = append(t.Headings, headings...)
}

func (t *Table) AddHeadingsS(headings ...string) {
	t.AddHeadings(stringSplitToInterfaceSplit(headings)...)
}

func (t *Table) AddRow(rowValues ...interface{}) {
	var thisRow []interface{}
	t.Rows = append(t.Rows, append(thisRow, rowValues...))
}

func stringSplitToInterfaceSplit(in []string) []interface{} {
	out := make([]interface{}, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}

func (t *Table) Print() {
	var headerFmt table.Formatter
	var columnFmt table.Formatter
	if shouldColor() {
		headerFmt = color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt = color.New(color.FgYellow).SprintfFunc()
	} else {
		headerFmt = fmt.Sprintf
		columnFmt = fmt.Sprintf
	}

	tbl := table.New(t.Headings...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, row := range t.Rows {
		tbl.AddRow(row...)
	}

	tbl.Print()
}
