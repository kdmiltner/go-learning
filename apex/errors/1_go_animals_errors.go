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

// Sound implementation.
func (a Animal) Sound() (string, error) {
	var (
		// Use the standard library errors package to create new errors.
		animalErr  = errors.New("the attribute `Animal` cannot be empty")
		noiseErr   = errors.New("the attribute `Noise` cannot be empty")
		wrappedErr error
	)

	if a.Animal == "" {
		// Use errors I created based on logic implemented to catch an error scenario.
		wrappedErr = animalErr
	}

	if a.Noise == "" {
		// Use errors I created based on logic implemented to catch an error scenario.
		// Since we want to catch and return all errors I wrap them using the errors.Join method
		// from the standard library.
		wrappedErr = errors.Join(noiseErr, wrappedErr)
	}

	output := fmt.Sprintf("The %s says %q", a.Animal, a.Noise)

	// Return wrappedErr and expect wherever this method is being called to handle the error.
	return output, wrappedErr
}

// MakeSound function.
func MakeSound(animal Sounder) error {
	output, err := animal.Sound()

	// Handle the error because we don't want to print the output if there is an error.
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}

func main() {
	a := Animal{Animal: "dog", Noise: "woof"}

	if err := MakeSound(a); err != nil {
		// panic is a special type of function that should be used sparingly as it will terminate
		// a run completely. Typically, you'll see a panic used in cases where our application
		// cannot function without the data. E.g. if you're starting up your app and as part of that startup
		// you initialize your database. If the database initialization fails you'd want to panic
		// as this is a necessary component for your application to run.
		panic(err)
	}
}
