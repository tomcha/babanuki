package main

import (
	"testing"
)

func TestPlay(t *testing.T) {
	m := new(Master)
	p1 := Player{
		name_:   "taro",
		myHand:  new(Hand),
		table_:  new(Table),
		master_: m,
	}
	p1.DealCrad(Card{SUIT_CLUB, 1})
	p1.DealCrad(Card{SUIT_CLUB, 2})
	p1.DealCrad(Card{SUIT_CLUB, 3})
	p2 := Player{
		name_:   "hanako",
		myHand:  new(Hand),
		table_:  new(Table),
		master_: m,
	}
	p2.DealCrad(Card{SUIT_HEART, 1})
	p2.DealCrad(Card{SUIT_HEART, 2})
	p2.DealCrad(Card{SUIT_HEART, 3})

	m.RegisterPlayer(p1)
	m.RegisterPlayer(p2)

	p1.Play(&p1.master_.players_[1])
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
		name_:   "taro",
		myHand:  h,
		table_:  new(Table),
		master_: new(Master),
	}
	sh := p.ShowHand()
	for i := 0; i < len(p.myHand.hand_); i++ {
		if sh.hand_[i] != cards[i] {
			return
		}
	}
	t.Fatal("test fale")
}

func TestShowHand2(t *testing.T) {
	m := new(Master)
	p1 := Player{
		name_:   "taro",
		myHand:  new(Hand),
		table_:  new(Table),
		master_: m,
	}
	p2 := Player{
		name_:   "taro",
		myHand:  new(Hand),
		table_:  new(Table),
		master_: m,
	}
	m.RegisterPlayer(p1)
	m.RegisterPlayer(p2)

	p1.myHand.hand_ = append(p1.myHand.hand_, Card{SUIT_CLUB, 1})
	sh := p1.ShowHand()
	if len(sh.hand_) != 1 {
		t.Fatal("fale test")
	}
}

func TestRecieveCard(t *testing.T) {
	p := Player{
		name_:   "taro",
		myHand:  new(Hand),
		table_:  new(Table),
		master_: new(Master),
	}
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
	p := Player{
		name_:  "taro",
		myHand: &Hand{},
		table_: &Table{},
	}
	c1 := Card{
		suit_:   SUIT_CLUB,
		number_: 10,
	}
	p.DealCrad(c1)
	if len(p.myHand.hand_) != 1 {
		t.Fatal("test fale")
	}
	c2 := Card{
		suit_:   SUIT_SPADE,
		number_: 11,
	}
	p.DealCrad(c2)
	if len(p.myHand.hand_) != 2 {
		t.Fatal("test fale")
	}
	c3 := Card{
		suit_:   SUIT_SPADE,
		number_: 10,
	}
	p.DealCrad(c3)
	if len(p.myHand.hand_) != 1 {
		t.Fatal("test fale")
	}
}

func TestString(t *testing.T) {
	p := Player{
		name_:   "taro",
		myHand:  new(Hand),
		table_:  new(Table),
		master_: new(Master),
	}
	if p.name_ != "taro" {
		t.Fatal("test fail")
	}
}
