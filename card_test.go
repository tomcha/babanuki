package main

import (
	"testing"
)

func TestCard(t *testing.T) {
	var s int = SUIT_DIAMOND
	var n int = 3

	c := Card{s, n}

	result, err := c.getNumber()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 3 {
		t.Fatal("faild test")
	}
}
