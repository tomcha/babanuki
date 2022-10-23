package main

import (
	"math/rand"
	"time"
)

type Hand struct {
	hand_ []Card
}

func (h *Hand) addCard(c Card) {
	h.hand_ = append(h.hand_, c)
}

func (h *Hand) pickCard() (c Card) {
	pickedCard := h.hand_[0]
	h.hand_ = h.hand_[1:]
	return pickedCard

}

func (h *Hand) shuffle() {
	size := len(h.hand_)
	for i := 0; i < size*2; i++ {
		rand.Seed(time.Now().UnixNano())
		pos := rand.Intn(size)
		if pos == 0 {
			c := h.hand_[0]
			h.hand_ = h.hand_[1:]
			h.hand_ = append(h.hand_, c)
		} else if pos != size {
			c := h.hand_[pos]
			newHand := append(h.hand_[0:pos], h.hand_[pos+1:size]...)
			h.hand_ = append(newHand, c)
		}
	}
}

func (h Hand) getNumberOfCard() int {
	return len(h.hand_)
}
