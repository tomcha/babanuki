package main

import (
	"testing"
)

func TestPlay(t *testing.T) {

}

func TestShowHand(t *testing.T) {
	h := new(Hand)
	cards := []Card{
		{SUIT_CLUB, 1},
		{SUIT_CLUB, 2},
		{SUIT_CLUB, 3},
		{SUIT_CLUB, 4},
		{SUIT_CLUB, 5},
		{SUIT_CLUB, 6},
		{SUIT_CLUB, 7},
		{SUIT_CLUB, 8},
		{SUIT_CLUB, 9},
		{SUIT_CLUB, 10},
		{SUIT_CLUB, 11},
		{SUIT_CLUB, 12},
		{SUIT_CLUB, 13},
	}
	cards2 := make([]Card, len(cards))
	copy(cards2, cards)
	h.hand_ = cards2
	p := Player{
		name_:  "taro",
		myHand: h,
		table_: new(Table),
	}
	p.ShowHand()
	if len(p.myHand.hand_) != len(cards) {
		t.Fatal("test fale")
	}
	for i := 0; i < len(p.myHand.hand_); i++ {
		if p.myHand.hand_[i] != cards[i] {
			return
		}
	}
	t.Fatal("test fale")
}

func TestRecieveCard(t *testing.T) {
	p := Player{name_: "taro", myHand: new(Hand), table_: new(Table)}
	c1 := Card{SUIT_SPADE, 10}
	c2 := Card{SUIT_DIAMOND, 10}
	p.RecieveCard(c1)
	if len(p.myHand.hand_) != 1 {
		t.Fatal("test fale")
	}
	p.RecieveCard(c2)
	if len(p.myHand.hand_) != 0 {
		t.Fatal("test fale")
	}
}

func TestDealCard(t *testing.T) {

}

func TestString(t *testing.T) {
	p := Player{"taro", new(Hand), new(Table)}
	if p.name_ != "taro" {
		t.Fatal("test fail")
	}
}
