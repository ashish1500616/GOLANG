package main

import (
	"fmt"
)

/* */
func main() {
	//----Declaring varaible---- .
	// ------
	var i int = 4
	fmt.Println(i)
	// ------
	var j int
	j = 5
	fmt.Println(j)
	// ------
	// Implicit initialization.
	auto_var := 5
	fmt.Println(auto_var)
	// ------
	println("----------------")
	//----WORKING WITH POINTERS-----
	var fn *string = new(string)
	*fn = "Ashish"
	fmt.Println(fn, *fn)

	lastName := "Verma"
	lnp := &lastName
	fmt.Println(lnp, *lnp)
	println("----------------")
	//----CONSTANT----
	const pi = 3.14
	fmt.Println(pi)
	//Cant change the value of pi and it should be defined while declaration .

	const c = 3

	//Case :  When the type is not defined for the variable c.
	fmt.Println(c + 3)   // Implicit identification of type for c.
	fmt.Println(c + 1.2) // Implicit identification of type for c.

	//if type is defined we have to do type casting.

	const d int = 3
	fmt.Println(float32(c) + 1.2)

	//----Using IOTA and Constants----
	// IOTA - Long chain of constants incremented by one .
	// IOTA - Resets within the constant block.

	//*****Collections*****
	//----Working with arr---- .
	println("----------------")
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	// ------

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	// ------
	// *. Slicing *. //
	slice := arr2[:] //included : excluded
	fmt.Println(slice)

	slice[1] = 4
	arr2[2] = 5
	fmt.Println(slice, arr2)

	//Slice can be think of as a pointer but it is not actually a pointer slice is pointing the underlying arr2.
	println("----------------")
	//Unbounded Array .
	un_a := []int{1, 2, 3}
	fmt.Println(un_a)

	un_a = append(un_a, 4, 5, 6)
	fmt.Println(un_a)

	// Resizing takes - O(n)

	s2 := un_a[1:]
	s3 := un_a[:2]
	s4 := un_a[1:2]
	fmt.Println(s2, s3, s4)
	println("----------------")
	//*. Maps .*
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 24
	fmt.Println(m["foo"])
	//** Delete Operation **
	delete(m, "foo")
	fmt.Println(m)

	//*. Structs .*
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)
	u2 := user{ID: 2, FirstName: "Arthur", LastName: "Dent"}
	fmt.Println(u2)

	//---- Creating Function and Methods . ----

	// ----
	// Loops .
	println("----------------")
	var k int
	for k < 5 {
		println(k)
		k++
		if k == 3 {
			break
		}
	}
	println("----------------")
	// Loop To Condition.
	//x := 0
	for x := 0; x < 5; x++ {
		println(x)
	}

	println("----------------")
	// UGLY Loop And Alternate.
	// for ; ;{}
	// for {}

	// -----------------
	// Looping Over a Collection .
	slice_loop :=
		[]int{1, 2, 3}
	for _, f := range slice_loop {
		println(f)
	}
	println("----------------")
	// Working With Panic
	// Hit A Condition Without which we cannot proceed .
	print("Panic : = p anic(Message)")

	//Writing Switch Statements .

	type HTTPRequest struct {
		Method string
	}

	r := HTTPRequest{Method: "HEAD"}

	switch r.Method {
	case "GET":
		println("Get request")
	case "DELETE":
		println("Delete request")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	default:
		println("Unhandled method")
	}

}
