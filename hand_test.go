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
	c3 := Card{SUIT_HEART, 12}

	h := Hand{}
	h.addCard(c1)
	h.addCard(c2)
	h.addCard(c3)

	if h.pickCard() != c1 {
		t.Fatal("test fale")
	}

	if h.hand_[0] != c2 || h.hand_[1] != c3 {
		t.Fatal("test fale")
	}
}
