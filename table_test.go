package main

import "testing"

func TestDisposeCard(t *testing.T) {
	table := Table{}
	c := Card{SUIT_SPADE, 10}
	table.disposeCard(c)
	if table.cards[len(cards)] != c {
		t.Fatal("test fale")
	}
}
