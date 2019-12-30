package main

import (
	"container/list"
	"fmt"
	"reflect"
)

// AnimalIntf - everyone likes animals
type AnimalIntf interface {
	TagID() int
}

// CatIntf is an animal
type CatIntf interface {
	AnimalIntf
	Meow()
}

// DogIntf is an animal
type DogIntf interface {
	AnimalIntf
	Bark()
}

// Shelter implements an animal shelter
type Shelter struct {
	animals *list.List
	cats    *list.List
	dogs    *list.List
}

// Animal implements AnimalIntf
type Animal struct {
	tagID int
}

// TagID returns the animal's tag ID
func (a *Animal) TagID() int {
	return a.tagID
}

// Cat implements CatIntf
type Cat struct {
	Animal
}

// Meow makes the sound of a dog
func (c *Cat) Meow() {
	fmt.Println(c.TagID(), "Meow! Meooow!")
}

// Dog implements DogIntf
type Dog struct {
	Animal
}

// Bark makes the sound of a dog
func (d *Dog) Bark() {
	fmt.Println(d.TagID(), "Woof! Woof!")
}

// New creates a new shelter
func New() *Shelter {
	s := new(Shelter)
	s.animals = list.New()
	s.cats = list.New()
	s.dogs = list.New()
	return s
}

// Enqueue adds the animal to the shelter
func (s *Shelter) Enqueue(a AnimalIntf) {
	s.animals.PushBack(a)
	switch v := a.(type) {
	case DogIntf:
		s.dogs.PushBack(a)
		v.Bark()
	case CatIntf:
		s.cats.PushBack(a)
		v.Meow()
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// DequeueAny removes the oldest animal - either cat or dog -
// from the shelter
func (s *Shelter) DequeueAny() AnimalIntf {
	var animal AnimalIntf
	for animal == nil && s.animals.Len() > 0 {
		a := s.animals.Front().Value.(AnimalIntf)
		s.animals.Remove(s.animals.Front())
		switch v := a.(type) {
		case DogIntf:
			// check if dog has not yet been given away
			if s.dogs.Front() != nil {
				dog := s.dogs.Front().Value.(DogIntf)
				if a == dog {
					// dog still there
					animal = dog
					s.dogs.Remove(s.dogs.Front())
				}
			}
		case CatIntf:
			// check if cat has not yet been given away
			if s.cats.Front() != nil {
				cat := s.cats.Front().Value.(CatIntf)
				if a == cat {
					// cat still there
					animal = cat
					s.cats.Remove(s.cats.Front())
				}
			}
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}
	return animal
}

// DequeueCat dequeues the oldest cat from the shelter
func (s *Shelter) DequeueCat() CatIntf {
	var cat CatIntf
	if s.cats.Len() > 0 {
		cat = s.cats.Front().Value.(CatIntf)
		s.cats.Remove(s.cats.Front())
	}
	return cat
}

// DequeueDog dequeues the oldest dog from the shelter
func (s *Shelter) DequeueDog() DogIntf {
	var dog DogIntf
	if s.dogs.Len() > 0 {
		dog = s.dogs.Front().Value.(DogIntf)
		s.dogs.Remove(s.dogs.Front())
	}
	return dog
}

func main() {
	shelter := New()

	d := new(Dog)
	d.tagID = 1
	shelter.Enqueue(d)
	c := new(Cat)
	c.tagID = 2
	shelter.Enqueue(c)

	c = new(Cat)
	c.tagID = 3
	shelter.Enqueue(c)
	d = new(Dog)
	d.tagID = 4
	shelter.Enqueue(d)

	a := shelter.DequeueAny()
	fmt.Println("Oldest animal (any)", reflect.TypeOf(a), "Tag ID", a.TagID())
	a = shelter.DequeueDog()
	fmt.Println("Oldest dog", reflect.TypeOf(a), "Tag ID", a.TagID())
	a = shelter.DequeueCat()
	fmt.Println("Oldest cat", reflect.TypeOf(a), "Tag ID", a.TagID())
	a = shelter.DequeueAny()
	fmt.Println("Oldest animal (any)", reflect.TypeOf(a), "Tag ID", a.TagID())
}
