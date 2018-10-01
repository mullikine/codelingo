package main

import "fmt"

func main() {
	var a ExportedStructA
	a.m = map[string]bool{
		"exhibit A": true,
	}

	fmt.Printf("%t\n", a.m["exhibit A"])

	// fmt.Printf("%s\n", ExportedConst)
}
