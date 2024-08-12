package main

import "fmt"

// Sounder3 interface has one method Sound that needs to be implemented for the
// contract to be satisfied.
type Sounder3 interface {
	Sound() string
}

// Animal2 struct represents the data we will use to implement the Sounder interface.
type Animal2 struct {
	Animal string
	Noise  string
}

// Sound implementation receives an Animal struct and returns a string, satisfying the Sounder interface.
func (a Animal2) Sound() string {
	return fmt.Sprintf("The %s says %q", a.Animal, a.Noise)
}

// Vehicle struct represents the data we will use to implement the Sounder interface.
type Vehicle struct {
	Vehicle string
	Make    string
	Model   string
	Noise   string
}

// Sound implementation receives a Vehicle struct and returns a string, satisfying the Sounder interface.
func (v Vehicle) Sound() string {
	if v.Make == "" && v.Model == "" {
		return fmt.Sprintf("The %s goes %q!", v.Vehicle, v.Noise)
	}

	return fmt.Sprintf("The %s goes %q and it's a %s %s!", v.Vehicle, v.Noise, v.Make, v.Model)
}

// MakeSound2 function
func MakeSound2(s Sounder3) {
	fmt.Println(s.Sound())
}

//func main() {
//	// Since a is of type Animal2 and Animal implements a method Sound it satisfies
//	// the interface and can be considered a Sounder.
//	a := Animal2{Animal: "dog", Noise: "woof"}
//	MakeSound2(a) // Output: The dog says 'woof'
//
//	// v is also a Sounder.
//	v := Vehicle{Vehicle: "train", Noise: "choo choo"}
//	MakeSound2(v) // Output: The train goes 'choo choo'!
//
//	v.Make = "Bullet"
//	v.Model = "3000"
//	MakeSound2(v) // Output: The train goes 'choo choo' and it's a Bullet 3000!
//}
