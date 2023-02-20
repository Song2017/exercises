package advance

import "fmt"

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (d Dog) Name() string {
	return d.name
}

func (d Dog) Category() string {
	return "cate dog"
}

func TestInterface() {
	fmt.Println("\n TestInterface")

	dog := Dog{name: "pupy"}
	var pet Pet = dog
	fmt.Println("pet name, dog name: ", pet.Name(), dog.Name())
	dog.name = "dog pupy"
	fmt.Println("pet name, dog name: ", pet.Name(), dog.Name())
	dog.SetName("monster")
	fmt.Println("pet name, dog name: ", pet.Name(), dog.Name())
	dog2 := dog
	dog.name = "dog name dog1"
	fmt.Println("dog2 name, dog name: ", dog2.Name(), dog.Name())

}
