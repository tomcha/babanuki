package main

import (
	//	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Hand struct {
	hand_ []Card
}

func (h Hand) String() string {
	var str string
	for _, c := range h.hand_ {
		str += c.String() + " "
	}
	str = strings.TrimRight(str, " ")
	return str
}

func (h *Hand) AddCard(c Card) {
	h.hand_ = append(h.hand_, c)
}

func (h *Hand) PickCard() (c Card) {
	pickedCard := h.hand_[0]
	h.hand_ = h.hand_[1:]
	return pickedCard

}

func (h *Hand) Shuffle() {
	size := len(h.hand_)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size*2; i++ {
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

func (h Hand) GetNumberOfCard() int {
	return len(h.hand_)
}

func (h *Hand) FindSameNumberCard() (sameCard []Card) {
	if len(h.hand_) == 0 {
		return sameCard
	}
	size := len(h.hand_)
	lc := h.hand_[size-1]
	for i := 0; i < size-1; i++ {
		if h.hand_[i].number_ == lc.number_ {
			sameCard = append(sameCard, h.hand_[i])
			sameCard = append(sameCard, lc)
			pickedHandFront := h.hand_[:i]
			pickedHandBack := h.hand_[i+1 : len(h.hand_)-1]
			//			h.hand_ = h.hand_[:i+(copy(h.hand_[i:], h.hand_[i+1:len(h.hand_)-2]))]
			h.hand_ = append(pickedHandFront, pickedHandBack...)
			return sameCard
		}
	}
	return sameCard
}
