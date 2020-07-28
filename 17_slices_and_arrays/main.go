package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	slice := arr[:]
	fmt.Println(slice)

	arr[0] = 42
	slice[2] = 27

	fmt.Println(arr, slice)
}
