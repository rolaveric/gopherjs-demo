package main

import (
	"github.com/rolaveric/gopherjs-demo/user"
	"github.com/rolaveric/gopherjs-demo/user/go/db"
	"fmt"
)

func main() {
	fmt.Println("Starting")

	fmt.Println("Registering the DB adapter");
	user.RegisterDB(db.DB{});

	fmt.Println("Getting all users");
	for _, v := range user.All() {
		fmt.Println(v)
	}

	fmt.Println("Adding a new user");
	user.New("Richard");
	fmt.Println(user.Get(6));

	fmt.Println("Update the name for a user");
	u := user.Get(2);
	u.Name = "Jason";
	user.Save(u);
	for _, v := range user.All() {
		fmt.Println(v)
	}

	fmt.Println("Done!")
}
