package main

import (
	"fmt"
	"log"

	"github.com/felipemaxplay/hello-go/math"
)

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

	resultado := math.Soma(1, 2)

	fmt.Println(resultado)

	res, err := math.MultMax10(1, 10)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)

	result := math.SumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)
}
