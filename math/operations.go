package math

import "errors"

// Soma return sum of x and y
func Soma(x int, y int) int {
	return x + y
}

// MultMax10 retorna a multiplicação de x por y
func MultMax10(x int, y int) (int, error) {
	res := x * y

	if y > 10 {
		return 0, errors.New("the multiplier cannot be greater than 10")
	}

	return res, nil
}
