package main

import (
	"testing"
)

func TestPlay(t *testing.T) {

}

func TesstShowHand(t *testing.T) {

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
