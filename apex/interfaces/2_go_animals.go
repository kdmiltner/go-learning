package main

import "fmt"

// Sounder interface
type Sounder interface {
	Sound() string
}

// Dog struct
type Dog struct{}

// Cat struct
type Cat struct{}

// Cow struct
type Cow struct{}

// Sound method implementation for Dog
func (d Dog) Sound() string {
	return "Woof!"
}

// Sound method implementation for Cat
func (c Cat) Sound() string {
	return "Meow!"
}

// Sound method implementation for Cow
func (c Cow) Sound() string {
	return "Moo!"
}

// makeSound function takes a Sounder interface
func makeSound(animal Sounder) {
	fmt.Println(animal.Sound())
}

//func main() {
//	dog := Dog{}
//	cat := Cat{}
//	cow := Cow{}
//
//	makeSound(dog) // Output: Woof!
//	makeSound(cat) // Output: Meow!
//	makeSound(cow) // Output: Moo!
//}
