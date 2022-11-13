package main

import (
	"testing"
)

func TestAddCard(t *testing.T) {
	c1 := Card{SUIT_DIAMOND, 10}
	c2 := Card{SUIT_CLUB, 3}

	h := Hand{}
	h.AddCard(c1)
	h.AddCard(c2)

	if h.hand_[1] != c2 {
		t.Fatal("test fale")
	}
}

func TestPickCard(t *testing.T) {
	c1 := Card{SUIT_DIAMOND, 10}
	c2 := Card{SUIT_CLUB, 3}
	c3 := Card{SUIT_HEART, 12}

	h := Hand{}
	h.AddCard(c1)
	h.AddCard(c2)
	h.AddCard(c3)

	if h.PickCard() != c1 {
		t.Fatal("test fale")
	}

	if h.hand_[0] != c2 || h.hand_[1] != c3 {
		t.Fatal("test fale")
	}
}

func TestShuffle(t *testing.T) {
	c1 := Card{SUIT_DIAMOND, 10}
	c2 := Card{SUIT_CLUB, 3}
	c3 := Card{SUIT_HEART, 12}

	h := Hand{}
	h.AddCard(c1)
	h.AddCard(c2)
	h.AddCard(c3)

	h.Shuffle()
	if len(h.hand_) != 3 {
		t.Fatal("test fale")
	}
}

func TestGetNumberOfCard(t *testing.T) {
	c1 := Card{SUIT_DIAMOND, 10}
	c2 := Card{SUIT_CLUB, 3}
	c3 := Card{SUIT_HEART, 12}

	h := Hand{}
	h.AddCard(c1)
	h.AddCard(c2)
	h.AddCard(c3)

	n := h.GetNumberOfCard()
	if n != len(h.hand_) {
		t.Fatal("test fale")
	}
}

func TestFindSameNumberCard(t *testing.T) {
	c1 := Card{SUIT_DIAMOND, 10}
	c2 := Card{SUIT_CLUB, 3}
	c3 := Card{SUIT_HEART, 12}

	h := Hand{}
	h.AddCard(c1)
	h.AddCard(c2)
	h.AddCard(c3)
	sc := h.FindSameNumberCard()
	if len(sc) != 0 {
		t.Fatal("test fale")
	}
	if len(h.hand_) != 3 {
		t.Fatal("test fale")
	}

	cc1 := Card{SUIT_DIAMOND, 10}
	cc2 := Card{SUIT_CLUB, 3}
	cc3 := Card{SUIT_HEART, 10}

	hh := Hand{}
	hh.AddCard(cc1)
	hh.AddCard(cc2)
	hh.AddCard(cc3)
	ssc := hh.FindSameNumberCard()
	if ssc[0] != cc1 || ssc[1] != cc3 {
		t.Fatal("test fale")
	}
	if hh.hand_[0] != cc2 {
		t.Fatal("test fale")
	}
}

func TestHandString(t *testing.T) {
	c1 := Card{SUIT_SPADE, 1}
	c2 := Card{SUIT_CLUB, 3}
	c3 := Card{SUIT_HEART, 12}
	c4 := Card{JOKER, 0}
	h := Hand{}
	h.AddCard(c1)
	h.AddCard(c2)
	h.AddCard(c3)
	h.AddCard(c4)
	str := h.String()

	if str != "SA C3 HQ JK" {
		t.Fatal("test fale")
	}
}
