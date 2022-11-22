package main

import (
	"testing"
)

func TestDeclareWin(t *testing.T) {
	m := new(Master)
	p := Player{
		name_:  "taro",
		myHand: new(Hand),
		table_: &Table{
			cards_: []Card{},
		},
	}
	m.players_ = append(m.players_, p)
	m.players_ = append(m.players_, p)
	m.players_ = append(m.players_, p)
	pn := len(m.players_)
	m.declareWin(p)
	if pn != len(m.players_)+1 {
		t.Fatal("test fale")
	}
}
