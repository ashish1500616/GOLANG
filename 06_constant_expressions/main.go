package main

const (
	first  = iota + 6
	second = 2 << iota
)

func main() {
	println(first, second)
}
