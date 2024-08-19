// package main
// goplayground link: https://go.dev/play/p/ghcb1JgAcxu
package main

import (
	"errors"
	"fmt"
)

// Sounder interface has one method Sound that needs to be implemented for the
// contract to be satisfied.
type Sounder interface {
	Sound() (string, error)
}

// Animal struct represents the data we will use to implement the Sounder interface.
type Animal struct {
	Animal string
	Noise  string
}

// Sound implementation receives an animal struct and returns a string.
func (a Animal) Sound() (string, error) {
	var (
		animalErr  = errors.New("the attribute `Animal` cannot be empty")
		noiseErr   = errors.New("the attribute `Noise` cannot be empty")
		wrappedErr error
	)

	if a.Animal == "" {
		wrappedErr = animalErr
	}

	if a.Noise == "" {
		wrappedErr = errors.Join(noiseErr, wrappedErr)
	}

	output := fmt.Sprintf("The %s says %q", a.Animal, a.Noise)

	return output, wrappedErr
}

// MakeSound function
func MakeSound(animal Sounder) error {
	output, err := animal.Sound()

	fmt.Println(output)
	return err
}

func main() {
	// Since a is of type Animal and Animal implements a method Sound it satisfies
	// the Sounder interface and can be considered a Sounder type.
	a := Animal{Animal: "dog", Noise: "woof"}

	if err := MakeSound(a); err != nil {
		panic(err)
	}
}
