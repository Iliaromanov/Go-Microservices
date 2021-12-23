package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:5000")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"First", "first item to add to database"}
	b := Item{"Second", "a second item"}
	c := Item{"Third", "third item"}

	// Add items to database
	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)

	// Retrieve a copy of the database
	client.Call("API.GetDB", "new db", &db)
	fmt.Println(db)

	// Using other methods
	client.Call("API.UpdateItem", Item{"Third", "updated body for the third item"}, &reply)
	fmt.Println(db)

	client.Call("API.DeleteItem", a, &reply)
	fmt.Println(db)

	client.Call("API.GetByTitle", "Second", &reply)
	fmt.Println(reply)
}