// package main
// goplayground link: https://go.dev/play/p/_AZgjw3iCUJ
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
// Vehicle doesn't need to be structured in the say way as Animal to implement the Sounder interface.
type Vehicle struct {
	Vehicle string
	Make    string
	Model   string
	Noise   string
}

// Speaker struct represents the data associated with an electronic speaker.
// Speaker does not have a Sound method, so it does not satisfy the Sounder interface.
type Speaker struct {
	Brand      string
	Size       float64
	MaxDecibel int
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

// MakeSound function can be called with an Animal or Vehicle since they implement Sounder.
func MakeSound(s Sounder) {
	fmt.Println(s.Sound())
}

func main() {
	// Since a is of type Animal and Animal implements a method Sound it satisfies
	// the Sounder interface and can be considered a Sounder.
	a := Animal{Animal: "dog", Noise: "woof"}
	MakeSound(a) // Output: The dog says 'woof'

	// v is also a Sounder since Vehicle implements a method Sound that satisfies the
	// Sounder interface.
	v := Vehicle{Vehicle: "train", Noise: "choo choo"}
	MakeSound(v) // Output: The train goes 'choo choo'!

	v.Make = "Bullet"
	v.Model = "3000"
	fmt.Println(v.Sound()) // Output: The train goes 'choo choo' and it's a Bullet 3000!

	// Since s is of type Speaker and Speaker does not have a Sound method it does
	// not implement or satisfy the Sounder interface and cannot be passed to MakeSound().
	//s := Speaker{Brand: "Bose", Size: 6.25, MaxDecibel: 90}
	//MakeSound(s)
}
