package main

import "fmt"

var ExportedVar string = "baz"
var unexportedVar string = "baz"

const ExportedConst string = "baz"
const unexportedConst string = "baz"

type ExportedStructA struct {
	m map[string]bool
}

type ExportedStructB struct {
	m map[string]bool
}

type unexportedStructA struct {
	m map[string]bool
}

func unexportedFunc() {
	var b unexportedStructA
	b.m = map[string]bool{
		"exhibit B": true,
	}

	fmt.Printf("%t\n", b.m["exhibit B"])
}

func ExportedFunc() {
	var b unexportedStructA
	b.m = map[string]bool{
		"exhibit C": true,
	}

	fmt.Printf("%t\n", b.m["exhibit C"])
}
