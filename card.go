package main

import "strconv"

const (
	JOKER = iota
	SUIT_SPADE
	SUIT_DIAMOND
	SUIT_CLUB
	SUIT_HEART
)

type Card struct {
	suit_   int
	number_ int
}

func (c Card) getNumber() (int, error) {
	return c.number_, nil
}

func (c Card) getSuit() (int, error) {
	return c.suit_, nil
}

func (c Card) String() string {
	var s string
	if c.suit_ == JOKER {
		return "JK"
	}
	switch c.suit_ {
	case SUIT_SPADE:
		s = "S"
	case SUIT_DIAMOND:
		s = "D"
	case SUIT_CLUB:
		s = "C"
	case SUIT_HEART:
		s = "H"
	}

	switch c.number_ {
	case 1:
		s += "A"
	case 11:
		s += "J"
	case 12:
		s += "Q"
	case 13:
		s += "K"
	default:
		s += strconv.Itoa(c.number_)
	}
	return s
}
