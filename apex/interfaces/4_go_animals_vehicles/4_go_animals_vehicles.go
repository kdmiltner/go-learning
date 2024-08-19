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

// Vehicle struct represents the data we will use to implement the Sounder interface.
type Vehicle struct {
	Vehicle string
	Make    string
	Model   string
	Noise   string
}

// Sound implementation receives an Animal struct and returns a string, satisfying the Sounder interface.
func (a Animal) Sound() string {
	return fmt.Sprintf("The %s says %q", a.Animal, a.Noise)
}

// Sound implementation receives a Vehicle struct and returns a string, satisfying the Sounder interface.
func (v Vehicle) Sound() string {
	if v.Make == "" && v.Model == "" {
		return fmt.Sprintf("The %s goes %q!", v.Vehicle, v.Noise)
	}

	return fmt.Sprintf("The %s goes %q and it's a %s %s!", v.Vehicle, v.Noise, v.Make, v.Model)
}

// MakeSound function
func MakeSound(s Sounder) {
	fmt.Println(s.Sound())
}

func main() {
	// Since a is of type Animal2 and Animal implements a method Sound it satisfies
	// the interface and can be considered a Sounder.
	a := Animal{Animal: "dog", Noise: "woof"}
	MakeSound(a) // Output: The dog says 'woof'

	// v is also a Sounder.
	v := Vehicle{Vehicle: "train", Noise: "choo choo"}
	MakeSound(v) // Output: The train goes 'choo choo'!

	v.Make = "Bullet"
	v.Model = "3000"
	MakeSound(v) // Output: The train goes 'choo choo' and it's a Bullet 3000!
}
