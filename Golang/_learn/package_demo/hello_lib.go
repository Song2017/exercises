package package_demo

import (
	"fmt"
)

// var block = "package"
// var block1 string

func TestPackage() {
	var test = "test string"
	test2 := "test2 string"
	fmt.Println(test, test2)

	ShowConst()

	TestOperator()

	_, r2, r3 := AddNumbers(12, 33, "the sum is: ")
	fmt.Println(r3, r2)

	block := "function"
	{
		block := "inner"
		fmt.Println("block3", block)
	}
	fmt.Println("block3", block)

	var err string
	fmt.Println(err)
	n, err := 0, ""
	fmt.Println(n, err)
}

func AddNumbers(a, b int, desc string) (int, int, string) {
	var result int = a + b

	return 0, result, desc
}

func ShowConst() {
	const a, b = "test", true
	// a = "modify"
	fmt.Print(a, b)
	fmt.Println("")
}

func TestOperator() {
	const a, b, c = 1, 2.5, 3
	var b1 bool
	b2, s1 := true, "string"

	fmt.Println(a+b+c, ", ", b1, ", ", b2)
	fmt.Println(string(a)+s1, string(a))
}
