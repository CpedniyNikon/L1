package tasks

import (
	"fmt"
	"math"
	"math/big"
)

func Multiply(a, b int64) (c big.Int) { // умножение больших чисел
	firstValue := big.NewInt(a)
	secondValue := big.NewInt(b)
	return *c.Mul(firstValue, secondValue)
}

func Difference(a, b int64) (c big.Int) { // вычитание больших чисел
	firstValue := big.NewInt(a)
	secondValue := big.NewInt(b)
	return *c.Sub(firstValue, secondValue)
}

func Sum(a, b int64) (c big.Int) { // сложение больших чисел
	firstValue := big.NewInt(a)
	secondValue := big.NewInt(b)
	return *c.Add(firstValue, secondValue)
}

func Divine(a, b int64) (c big.Int) { // деление больших чисел нацело
	firstValue := big.NewInt(a)
	secondValue := big.NewInt(b)
	return *c.Div(firstValue, secondValue)
}

func Task22() {
	a := math.Pow(2, 20)
	b := math.Pow(2, 21)

	result := Multiply(int64(a), int64(b))
	fmt.Println(result.String())
}
