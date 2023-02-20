package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func New(name string) Dog {
	return Dog{name}
}

func (d *Dog) SetName(n string) {
	d.name = n
}

func main() {
	//  Go 语言中的哪些值是不可寻址的
	fmt.Println("pointer")
	// 1. 不可变的值
	const num = 123
	// &num
	// 123
	// 2. 临时结果
	// &(123 + 123)
	str := "asdf"
	// &str[0]
	str2 := str[0]
	_ = &str2 // 对切片字面量的索引结果值虽然也属于临时结果，但却是可寻址的。
	//_ = &([3]int{1, 2, 3}[0]) // 对数组字面量的索引结果值不可寻址。
	//_ = &([3]int{1, 2, 3}[0:2]) // 对数组字面量的切片结果值不可寻址。
	_ = &([]int{1, 2, 3}[0]) // 对切片字面量的索引结果值却是可寻址的。
	// _ = &([]int{1, 2, 3}[0:2]) // 对切片字面量的切片结果值不可寻址。
	//_ = &(map[int]string{1: "a"}[0]) // 对字典字面量的索引结果值不可寻址。 字典中的每个键 - 元素对的存储位置都可能会变化
	dog := Dog{"little pig"}
	_ = dog
	//_ = &(dog.Name) // 标识符代表的函数不可寻址。
	//_ = &(dog.Name()) // 对方法的调用结果值不可寻址。
	// New("puppy").SetName("monster") // 结构体字面量的字段不可寻址。

	// 通过程序内部数据的内存地址查找变量
	dogP := &dog
	dogPtr := uintptr(unsafe.Pointer(dogP))
	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Println(*nameP)
}
