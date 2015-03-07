package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/rolaveric/gopherjs-demo/user"
	"github.com/rolaveric/gopherjs-demo/user/js/db"
)

// Starting point for compiling JS code
func main() {
	js.Global.Set("user", map[string]interface{}{
		"registerDB": RegisterDBJS,
		"new":        user.New,
		"get":        user.Get,
		"all":        user.All,
		"save":       SaveJS,
	})
}

// Takes a DB adapter written in Javascript and wraps it as a DB interface
func RegisterDBJS(o *js.Object) {
	user.RegisterDB(db.JSDB{o})
}

// Takes a JS object and wraps it as a User struct
func SaveJS(o *js.Object) {
	user.Save(&user.User{o.Get("Name").String(), o.Get("ID").Int()})
}
