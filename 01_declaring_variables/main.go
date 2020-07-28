package main

func main() {
	var i int
	i = 42
	println(i)

	var f float32 = 3.14
	println(f)

	firstName := "Arthur"
	println(firstName)

	c := complex(3, 4)
	println(c)

	a, b := real(c), imag(c)
	println(a, b)
}
