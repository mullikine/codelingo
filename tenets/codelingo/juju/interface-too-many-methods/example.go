// Package mypack This package serves as example code for the tenet `interface-too-many-methods`
package mypack

type smallInterface interface {
	surfaceArea() float64
	volume() float64
}

type excessiveInterface interface {
	perimeter() float64
	surfaceArea() float64
	volume() float64
	diameter() float64
	radius() float64
	circumference() float64
}
