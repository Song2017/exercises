package advance

import (
	"errors"
	"fmt"
)

func TestFunction() {
	fmt.Println("Function Test")

	var p Printer
	p = printStd
	fmt.Println(p("test"))

	r, _ := Calc(3, 4, add)
	fmt.Println("advance Function Calc add", r)
	mul := func(x, y int) (int, error) { return x * y, nil }
	mr, _ := Calc(3, 4, mul)
	fmt.Println("advance Function Calc mul", mr)

	add2 := genCalculator(add)
	result, err := add2(3, 4)
	fmt.Printf("return cal The result: %d (error: %v)\n", result, err)

	// parameter effect: shadow copy
	FunctionParameters()
}

type Printer func(contents string) (n int, err error)

func printStd(cont string) (bytesLen int, err error) {
	return fmt.Println(cont)
}

type operate func(x, y int) (int, error)

func Calc(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("test")
	}
	return op(x, y)
}

func add(x, y int) (int, error) {
	return x + y, nil
}

type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("return fun error")
		}
		return op(x, y)
	}
}

func FunctionParameters() {
	modify1 := func(arr [3]string) [3]string { arr[2] = "test"; return arr }
	arr := [3]string{"f", "s", "t"}
	arr2 := modify1(arr)
	fmt.Println("arr no effect", arr, arr2)
	modify2 := func(arr [2][]string) [2][]string { arr[1] = []string{"test"}; return arr }
	slice := [2][]string{{"111", "11111"}, {"2222", "22222"}}
	slice2 := modify2(slice)
	fmt.Println("arr[slice], arr ele no effect", slice, slice2)
	modify3 := func(arr [2][]string) [2][]string { arr[1][1] = "test"; return arr }
	slice3 := modify3(slice)
	fmt.Println("arr[slice], slice ele effect", slice, slice3)
}
