/*
	Open Closed Principle says that a class/function should only be open to extension but closed to modification. Basically,
	we should be able to use the existing class/function for new functionality without having to modify the said class/function.
*/

package main

import (
	"fmt"
	"log"
	"math"
)

/*
	-----------------------Bad Example--------------------------------
	The below AreaCalculator struct and the CalculateArea method are -
	1. Calculating the area of a shape
	2. Adding the area of a shape
	3. Returning the total area

	But, if we want to add a new shape for example a Triangle, we need to modify the CalculateArea method because to calculate the area of a triangle,
	we need to know the base and height of the triangle. This violates the Open Closed Principle.
*/

type AreaCalculator struct{}

func (ac *AreaCalculator) CalculateArea(shapes []interface{}) float64 {
	var area float64

	for _, shape := range shapes {
		switch s := shape.(type) {
		case Rectangle:
			area += s.Width * s.Height
		case Circle:
			area += math.Pi * s.Radius * s.Radius
		default:
			log.Println("Unknown shape")
		}
	}

	return area
}

/*
	-----------------------Good Example--------------------------------
	We will refactore the above code such that we can add new shapes without modifying the CalculateArea method.
	In this, what we will do is -
	1. Create a Shape interface that will have a method Area() float64
	2. Create structs for the shapes that implement the Shape interface and return the area of the shape
	3. Create a method that will calculate the total area of the shapes
	4. If we want to add a new shape, we can simply create a new struct that implements the Shape interface and return the area of the shape

	Now, we can add a new shape without modifying the CalculateArea method.
*/

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

type ShapeAreaCalculator struct{}

func (sac *ShapeAreaCalculator) CalculateArea(shapes []Shape) float64 {
	var totalArea float64

	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

func OCP_Demo() {
	fmt.Println("Open Closed Principle")

	shapes := []Shape{
		Rectangle{Width: 10, Height: 20},
		Circle{Radius: 5},
		Triangle{Base: 10, Height: 20},
	}

	areaCalculator := &ShapeAreaCalculator{}
	totalArea := areaCalculator.CalculateArea(shapes)
	fmt.Printf("Total Area - %f\n", totalArea)
	fmt.Println()
}
