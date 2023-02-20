package advance

import (
	"fmt"
	"unsafe"
)

type AnimalCategory struct {
	kingdom string
	phylum  string
	class   string

	species string
}

func (a AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s", a.kingdom, a.phylum, a.class,
		a.species)
}

type Animal struct {
	name string
	AnimalCategory
}

func (a Animal) Category() string {
	return fmt.Sprintf("cate %s, name %s", a.AnimalCategory.String(), a.name)
}
func (a Animal) String() string {
	return fmt.Sprintf("name %s", a.name)
}
func (a *Animal) SetName(name string) {
	a.name = name
}

func TestStruct() {
	fmt.Println("\n TestStruct")
	ac := AnimalCategory{species: "cat"}
	fmt.Println(ac)
	a := Animal{name: "nico", AnimalCategory: AnimalCategory{species: "cat"}}
	fmt.Println(a, a.String(), a.Category())

	ps := struct{}{}
	fmt.Println("uintptr(unsafe.Pointer(&struct{}{}))", ps, uintptr(unsafe.Pointer(&ps)), unsafe.Sizeof(struct{}{}))
	ps2 := struct{}{}
	fmt.Println("uintptr(unsafe.Pointer(&struct{}{}))", ps2, uintptr(unsafe.Pointer(&ps2)))
}
