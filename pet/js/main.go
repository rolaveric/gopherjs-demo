package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/rolaveric/gopherjs-demo/pet"
)

func main() {
	js.Global.Set("pet", map[string]interface{}{
		"New": pet.New,
	})
}
