// package main
// goplaygound link: https://go.dev/play/p/KdbtkQTjixa
package main

import "fmt"

// Sounder interface has one method Sound that needs to be implemented for the
// contract to be satisfied.
type Sounder interface {
	Sound() string
}

// Animal struct represents the data we will use to implement the Sounder interface.
type Animal struct {
	Animal string
	Noise  string
}

// Sound implementation receives an animal struct and returns a string.
func (a Animal) Sound() string {
	return fmt.Sprintf("The %s says %q", a.Animal, a.Noise)
}

// MakeSound function
func MakeSound(animal Sounder) {
	fmt.Println(animal.Sound())
}

func main() {
	// Since a is of type Animal and Animal implements a method Sound it satisfies
	// the Sounder interface and can be considered a Sounder type.
	a := Animal{Animal: "dog", Noise: "woof"}
	b := Animal{Animal: "cat", Noise: "meow"}
	c := Animal{Animal: "cow", Noise: "moo"}
	MakeSound(a) // Output: The dog says 'woof'
	fmt.Println(b.Sound())
	MakeSound(c)
}
