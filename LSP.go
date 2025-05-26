/*
	Liskov Substitution Princple says that if you have a parent type and a child type,
	you should be able to use the child anywhere you use the parent without breaking anything.
*/

package main

import "fmt"

/*
	-----------------------Bad Example--------------------------------
	In the below example, we have a FlyingBirdInterface that has a Fly method.
	But, we are trying to use it with a PenguinBird struct which is not a flying bird.
	So, this violates the Liskov Substitution Principle.
*/

type FlyingBirdInterface interface {
	Fly() string
}

type SparrowBird struct{}

func (s SparrowBird) Fly() string {
	return "Sparrow is flying"
}

type PenguinBird struct{}

func (p PenguinBird) Fly() string {
	return "Penguin is swimming" // This is wrong because a penguin cannot fly
}

/*
	-----------------------Good Example--------------------------------
	In the below example, we have a Bird interface that has a Move and Eat method.
	We have a FlyingBird interface that has a Bird interface and a Fly method.
	We have a SwimmingBird interface that has a Bird interface and a Swim method.
	We have a Sparrow struct that implements the FlyingBird interface.
	We have a Penguin struct that implements the SwimmingBird interface.

	Now, we can use the Sparrow and Penguin structs anywhere we use the Bird interface.
*/

type Bird interface {
	Move() string
	Eat() string
}

type FlyingBird interface {
	Bird
	Fly() string
}

type SwimmingBird interface {
	Bird
	Swim() string
}

type Sparrow struct{}

func (s Sparrow) Move() string {
	return "Sparrow is moving"
}

func (s Sparrow) Eat() string {
	return "Sparrow is eating"
}

func (s Sparrow) Fly() string {
	return "Sparrow is flying"
}

type Penguin struct{}

func (p Penguin) Move() string {
	return "Penguin is moving"
}

func (p Penguin) Eat() string {
	return "Penguin is eating"
}

func (p Penguin) Swim() string {
	return "Penguin is swimming"
}

func LSP_Demo() {
	fmt.Println("Liskov Substitution Principle")

	sparrow := Sparrow{}
	penguin := Penguin{}

	sparrow.Move()
	sparrow.Eat()
	sparrow.Fly()

	penguin.Move()
	penguin.Eat()
	penguin.Swim()

	fmt.Println()
}
