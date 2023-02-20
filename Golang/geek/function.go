package main_function

import (
	"errors"
	"fmt"
)

type Printer func(content string) (n int, err error)

func printToStd(c string) (bytes int, err error) {
	return fmt.Println(c)
}

type operate func(x, y int) (n int)

func calculate(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("op is invalid")
	}
	return op(x, y), nil
}

type calculateFunc func(x, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("op is invalid")
		}
		return op(x, y), nil
	}
}

func main() {
	var p Printer
	p = printToStd
	n, _ := p("test")
	fmt.Println(n)
	// 高阶函数可以满足下面的两个条件：
	// 1. 接受其他的函数作为参数传入；
	op1 := func(x, y int) int {
		return x + y
	}
	n2, _ := calculate(1, 23, op1)
	fmt.Println("func as parameters", n2)
	// 2. 把其他的函数作为结果返回。
	add := genCalculator(op1)
	n3, _ := add(12, 2)
	fmt.Println("func as response parameters", n3)
	// 传入函数的那些参数值
	array1 := [3][]string{
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
	}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1)
}

func modifyArray(a [3][]string) [3][]string {
	a[1][1] = "modifyArray"
	a[1] = []string{"x"}
	return a
}
