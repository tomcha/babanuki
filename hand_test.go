package main

import (
	"testing"
)

func TestAddCard(t *testing.T) {
	c1 := Card{SUIT_DIAMOND, 10}
	c2 := Card{SUIT_CLUB, 3}

	h := Hand{}
	h.addCard(c1)
	h.addCard(c2)

	if h.hand_[1] != c2 {
		t.Fatal("test fale")
	}
}

func TestPickCard(t *testing.T) {
	c1 := Card{SUIT_DIAMOND, 10}
	c2 := Card{SUIT_CLUB, 3}

	h := Hand{}
	h.addCard(c1)
	h.addCard(c2)

	if h.pickCard() != c1 {
		t.Fatal("test fale")
	}
}
