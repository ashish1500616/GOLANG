package main

import (
	"fmt"
)

func main() {
	var firstName *string = new(string)

	*firstName = "Arthur"
	fmt.Println(*firstName)
}
