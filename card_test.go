package main

import (
	"testing"
)

func TestGetNumber(t *testing.T) {
	var s int = SUIT_DIAMOND
	var n int = 3
	c := Card{s, n}

	result, err := c.getNumber()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 3 {
		t.Fatal("failed test")
	}

	s = SUIT_DIAMOND
	n = 10
	c = Card{s, n}

	result, err = c.getNumber()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 10 {
		t.Fatal("failed test")
	}
}

func TestGetSuit(t *testing.T) {
	var s int = SUIT_HEART
	var n int = 10
	c := Card{s, n}

	result, err := c.getSuit()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != SUIT_HEART {
		t.Fatal("failed test")
	}

	s = SUIT_CLUB
	n = 5
	c = Card{s, n}
	result, err = c.getSuit()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != SUIT_CLUB {
		t.Fatal("failed test")
	}
}

func TestString(t *testing.T) {
	var s int = SUIT_SPADE
	var n int = 1
	var r string
	c := Card{s, n}
	r = c.String()
	if r != "SA" {
		t.Fatal("failed test")
	}

	s = SUIT_HEART
	n = 11
	c = Card{s, n}
	r = c.String()
	if r != "HJ" {
		t.Fatal("failed test")
	}

	s = SUIT_DIAMOND
	n = 10
	c = Card{s, n}
	r = c.String()
	if r != "D10" {
		t.Fatal("failed test")
	}
}
