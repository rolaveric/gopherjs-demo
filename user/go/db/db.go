// Go implementation of user.DB interfaces
package db

import (
	"github.com/rolaveric/gopherjs-demo/user"
	"strconv"
	"strings"
)

// Hard coded rows
var rows []*DBRow = []*DBRow{
	&DBRow{[]string{"Jane", "1"}},
	&DBRow{[]string{"John", "2"}},
	&DBRow{[]string{"Sarah", "3"}},
	&DBRow{[]string{"Steve", "4"}},
	&DBRow{[]string{"Jess", "5"}},
}

// Implements the user.DBRow interface
type DBRow struct {
	Columns []string
}

func (row *DBRow) GetInt(colnum int) int {
	i, err := strconv.Atoi(row.Columns[colnum])
	if err != nil {
		panic(err)
	}
	return i
}

func (row *DBRow) GetString(colnum int) string {
	return row.Columns[colnum]
}

// Implements the user.DBResult interface
type DBResult struct {
	Rows []*DBRow
	Index int
}

func (result *DBResult) NextRow() user.DBRow {
	v := result.Rows[result.Index]
	result.Index++
	return v
}

func (result *DBResult) RowCount() int {
	return len(result.Rows)
}

// A new type that implements user.DB by using a Javascript object
type DB struct {}

// Uses the 'Query()' function from the JS object and turns the result into a DBResult
func (db DB) Query(query string, params ...interface{}) user.DBResult {
	switch {
	case strings.HasPrefix(query, "UPDATE"):
		rows[params[1].(int) - 1].Columns[0] = params[0].(string)
	case strings.HasPrefix(query, "INSERT"):
		rows = append(rows, &DBRow{[]string{params[0].(string), strconv.Itoa(len(rows) + 1)}})
	case query == "SELECT @@IDENTITY":
		row := &DBRow{[]string{strconv.Itoa(len(rows))}}
		return &DBResult{[]*DBRow{row}, 0}
	case strings.Contains(query, "WHERE"):
		return &DBResult{[]*DBRow{rows[params[0].(int) - 1]}, 0}
	case strings.HasPrefix(query, "SELECT"):
		return &DBResult{rows, 0}
	}
	return nil
}
