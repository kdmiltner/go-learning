package main

import "fmt"

// Sounder2 interface has one method Sound that needs to be implemented for the
// contract to be satisfied.
type Sounder2 interface {
	Sound() string
}

// Animal struct represents the data we will use to implement the Sounder interface.
type Animal struct {
	Type  string
	Noise string
}

// Sound implementation receives an animal struct and returns a string.
func (a Animal) Sound() string {
	return fmt.Sprintf("The %s says %q", a.Type, a.Noise)
}

// MakeSound function
func MakeSound(animal Sounder2) {
	fmt.Println(animal.Sound())
}

//func main() {
//	// Since a is of type Animal and Animal implements a method Sound it satisfies
//	// the Sounder interface and can be considered a Sounder type.
//	a := Animal{Type: "dog", Noise: "woof"}
//	MakeSound(a) // Output: The dog says 'woof'
//}
