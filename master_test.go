package main

import (
	"fmt"
	"testing"
)

func TestPrepareGame(t *testing.T) {
	m := new(Master)
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
		{SUIT_HEART, 1},
		{SUIT_HEART, 2},
		{SUIT_HEART, 3},
		{SUIT_HEART, 4},
		{SUIT_HEART, 5},
		{SUIT_HEART, 6},
		{SUIT_HEART, 7},
		{SUIT_HEART, 8},
		{SUIT_HEART, 9},
		{SUIT_HEART, 10},
		{SUIT_HEART, 11},
		{SUIT_HEART, 12},
		{SUIT_HEART, 13},
		{SUIT_SPADE, 1},
		{SUIT_SPADE, 2},
		{SUIT_SPADE, 3},
		{SUIT_SPADE, 4},
		{SUIT_SPADE, 5},
		{SUIT_SPADE, 6},
		{SUIT_SPADE, 7},
		{SUIT_SPADE, 8},
		{SUIT_SPADE, 9},
		{SUIT_SPADE, 10},
		{SUIT_SPADE, 11},
		{SUIT_SPADE, 12},
		{SUIT_SPADE, 13},
		{SUIT_DIAMOND, 1},
		{SUIT_DIAMOND, 2},
		{SUIT_DIAMOND, 3},
		{SUIT_DIAMOND, 4},
		{SUIT_DIAMOND, 5},
		{SUIT_DIAMOND, 6},
		{SUIT_DIAMOND, 7},
		{SUIT_DIAMOND, 8},
		{SUIT_DIAMOND, 9},
		{SUIT_DIAMOND, 10},
		{SUIT_DIAMOND, 11},
		{SUIT_DIAMOND, 12},
		{SUIT_DIAMOND, 13},
		{JOKER, 0},
	}
	h.hand_ = cards

	p1 := Player{
		name_:   "taro",
		myHand:  &Hand{},
		table_:  &Table{},
		master_: m,
	}
	p2 := Player{
		name_:   "hanako",
		myHand:  &Hand{},
		table_:  &Table{},
		master_: m,
	}
	p3 := Player{
		name_:   "jiro",
		myHand:  &Hand{},
		table_:  &Table{},
		master_: m,
	}
	m.RegisterPlayer(p1)
	m.RegisterPlayer(p2)
	m.RegisterPlayer(p3)
	m.PrepareGame(h)

	fmt.Println("p1: ", p1.myHand.hand_)
	fmt.Println("p2: ",p2.myHand.hand_)
	fmt.Println("p3: ",p3.myHand.hand_)
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
