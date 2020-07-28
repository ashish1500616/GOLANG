package main

import "fmt"

func main() {

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	u := user{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	fmt.Println(u)
}
