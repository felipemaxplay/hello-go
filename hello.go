package main

import "fmt"

func main() {
	fmt.Printf("Hello, world\n")

	a := 25
	b := "world"
	c := 1.921
	d := false

	fmt.Printf("%v -> %T\n", a, a)
	fmt.Printf("%v -> %T\n", b, b)
	fmt.Printf("%v -> %T\n", c, c)
	fmt.Printf("%v -> %T\n", d, d)
}
