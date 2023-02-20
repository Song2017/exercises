package main_struct_interface

import (
	"fmt"
	"reflect"
)

// struct
type AnimalCategory struct {
	genus   string
	species string
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("AnimalCategory genus %s, species %s", ac.genus, ac.species)
}

type Animal struct {
	Name           string
	AnimalCategory // 字段的声明中只有字段的类型名而没有字段的名称, 嵌入字段的类型既是类型也是名称
}

func (a Animal) Category() string {
	return a.AnimalCategory.String()
}

func (a Animal) String() string {
	return fmt.Sprintf("Animal category %s, Name %s", a.AnimalCategory.String(), a.Name)
}

type Cat struct {
	CatName string
	Animal  // 多层嵌入
}

func (a Cat) String() string {
	return fmt.Sprintf("Cat category %s, CatName %s", a.Animal.String(), a.CatName)
}

// interface
type Pet interface {
	SetName(name string)
	Name() string
}

type Dog struct {
	name string
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func (d Dog) Name() string {
	return d.name
}

func main() {
	fmt.Println("struct")
	ac := AnimalCategory{species: "cat"}
	fmt.Println(ac)
	a := Animal{Name: "American Shorthair"}
	fmt.Println(a)
	c := Cat{CatName: "kitty", Animal: a}
	fmt.Println(c)
	fmt.Println("interface")

	d := Dog{name: "puppy"}
	fmt.Println(d)
	fmt.Println(&d)
	pet := &d
	fmt.Println(pet.Name())
	fmt.Printf("The type of pet is %s.\n", reflect.TypeOf(pet).String())

	_, ok := interface{}(d).(Dog)
	fmt.Println("dog implement interface Dog:", ok)
	_, ok2 := interface{}(&d).(Pet)
	fmt.Println("dog implement interface Pet:", ok2)

	var d1 *Dog
	d2 := d1
	var p2 Pet = d2
	fmt.Println(p2 == nil)
	fmt.Printf("The type of d1 is %T.\n", d1)
	fmt.Printf("The type of pet is %T.\n", p2)

	dog := Dog{name: "dog1"}
	dog2 := dog
	var pet2 Pet = &dog2
	dog2.SetName("dog 2")
	fmt.Printf("Dog name %s.\n", dog.Name())
	fmt.Printf("Pet name %s.\n", pet2.Name())
}
