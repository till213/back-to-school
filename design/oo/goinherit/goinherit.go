package main

import (
	"fmt"
)

type animal struct {
}

func (a *animal) makeNoise() string {
	return "..."
}

type dog struct {
	animal
}

func (d *dog) makeNoise() string {
	return "Woof!"
}

type saintBernard struct {
	dog
}

func (s *saintBernard) makeNoise() string {
	return "Wuff! Wuff!"
}

func main() {
	a := new(animal)
	d := new(dog)
	s := new(saintBernard)

	fmt.Println("An animal makes:", a.makeNoise())
	fmt.Println("A dog makes:", d.makeNoise())
	fmt.Println("A Saint Bernard makes:", s.makeNoise())

	// This does not work (in Go): even though a1 seems
	// to be a "base class" of s it cannot be assigned
	// an s (not even with type-casting)
	// -> polyphormism is implementeing using interfaces
	//var a1 *animal
	//a1 = *animal(s)
}
