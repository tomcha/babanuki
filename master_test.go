package main

import (
	"testing"
)

func TestPrepareGame(t *testing.T) {
	m := new(Master)
	m.PrepareGame()
}

func TestStartGame(t *testing.T) {

}

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
	m.DeclareWin(p)
	if pn != len(m.players_)+1 {
		t.Fatal("test fale")
	}
}

func TestRegisterPlayer(t *testing.T) {
	m := new(Master)
	p1 := Player{
		name_:  "taro",
		myHand: &Hand{},
		table_: &Table{},
	}
	p2 := Player{
		name_:  "hanako",
		myHand: &Hand{},
		table_: &Table{},
	}
	m.RegisterPlayer(p1)
	m.RegisterPlayer(p2)
	if len(m.players_) != 2 {
		t.Fatal("test fatal")
	}
}
