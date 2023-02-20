package advance

import (
	"errors"
	"fmt"
)

func TestPRD() {
	fmt.Println("\n TestPRD")
	testPanic()
	testRecover()
	testDefer()
}

func testDefer() {
	fmt.Println("testDefer")
	defer func() {
		recover()
	}()
	defer fmt.Println("defer begin")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("defer end")
	defer func() {
		panic("panic in defer")
	}()

}

func testRecover() {
	fmt.Println("testRecover")

	defer func() {
		fmt.Println("Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()
	panic(errors.New("something wrong"))
	fmt.Println("Exit testRecover")
}

func testPanic() {
	fmt.Println("testPanic")
	// panic("test panic") // panic: test panic
	// panic(fmt.Sprintf("test panic")) // panic: test panic

}
