package main

import "testing"

func TestDisposeCard(t *testing.T) {
	table := Table{}
	table.cards_ = []Card{}
	c := Card{SUIT_SPADE, 10}
	table.DisposeCard(c)
	if table.cards_[len(table.cards_)-1] != c {
		t.Fatal("test fale")
	}
}

func TestDisposeCard2(t *testing.T) {
	table := Table{}
	table.cards_ = []Card{
		{SUIT_DIAMOND, 1},
		{SUIT_CLUB, 3},
	}
	c := Card{SUIT_SPADE, 7}
	table.DisposeCard(c)
	if table.cards_[len(table.cards_)-1] != c {
		t.Fatal("test fale")
	}
}
