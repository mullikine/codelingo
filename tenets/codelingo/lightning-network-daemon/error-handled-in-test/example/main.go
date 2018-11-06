package main

import "fmt"

import (
	"testing"
)

func SayA() error {
	return "I'm a ", nil
}

func SayB() error {
	return "I'm b ", nil
}
