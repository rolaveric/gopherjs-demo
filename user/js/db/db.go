// Contains types for implementing the DB, DBResult and DBRow interfaces using JS objects
// Assumes that the JS version of 'Query()' returns a matrix of rows and columns.
package db

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/rolaveric/gopherjs-demo/user"
)

// Implements the user.DBRow interface
type JSDBRow struct {
	O *js.Object // An array of column values
}

func (jsrow *JSDBRow) GetInt(colnum int) int {
	return jsrow.O.Index(colnum).Int()
}

func (jsrow *JSDBRow) GetString(colnum int) string {
	return jsrow.O.Index(colnum).String()
}

// Implements the user.DBResult interface
type JSDBResult struct {
	O *js.Object // An array of rows
	Index int // The current row index
}

func (jsresult *JSDBResult) NextRow() user.DBRow {
	v := &JSDBRow{jsresult.O.Index(jsresult.Index)}
	jsresult.Index++
	return v
}

func (jsresult *JSDBResult) RowCount() int {
	return jsresult.O.Length()
}

// A new type that implements user.DB by using a Javascript object
type JSDB struct {
	O *js.Object
}

// Uses the 'Query()' function from the JS object and turns the result into a DBResult
func (jsdb JSDB) Query(query string, params ...interface{}) user.DBResult {
	return &JSDBResult{jsdb.O.Call("Query", query, params), 0}
}

