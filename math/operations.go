package math

import "errors"

// Soma return sum of x and y
func Soma(x int, y int) (res int) {
	res = x + y
	return
}

// SumAll number in parameter
func SumAll(x ...int) int {
	result := 0
	for _, v := range x {
		result += v
	}
	return result
}

// MultMax10 retorna a multiplicação de x por y
func MultMax10(x int, y int) (int, error) {
	res := x * y

	if y > 10 {
		return 0, errors.New("the multiplier cannot be greater than 10")
	}

	return res, nil
}
