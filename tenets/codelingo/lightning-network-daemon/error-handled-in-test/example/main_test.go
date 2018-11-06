package main

import (
	"testing"
)

func TestSayA() {
	err := SayA()
}

func TestSayB() {
	err := SayB()
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
}
