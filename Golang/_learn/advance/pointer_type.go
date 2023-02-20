package advance

import (
	"fmt"
	"unsafe"
)

type Cat struct {
	name string
}

func (c *Cat) SetName(n string) {
	c.name = n
}

func New(name string) Cat {
	return Cat{name: name}
}

func TestPointer() {
	fmt.Println("\nbegin TestPointer")
	cat := Cat{name: "nico"}
	fmt.Println("&cat", &cat)

	// 由于New函数的调用结果值是不可寻址的，所以无法对它进行取址操作
	// New("nico2").SetName("nico2 name") //cannot call pointer method SetName on Cat
	cat2 := New("nico2")
	cat2.SetName("nico2 name")
	fmt.Println(cat2.name)

	// unsafe.Pointer是像*Dog类型的值这样的指针值和uintptr值之间的桥梁
	cat3 := Cat{name: "cat3"}
	catP := &cat3
	catPtr := uintptr(unsafe.Pointer(catP))
	// catPtr2 := uintptr(catP)
	fmt.Println(cat3, catP, catPtr)
	namePtr := catPtr + unsafe.Offsetof(catP.name)
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Println(namePtr, *nameP)
}
