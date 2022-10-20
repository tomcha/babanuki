package main

type Hand struct {
	hand_ []Card
}

func (h *Hand) addCard(c Card) {
	h.hand_ = append(h.hand_, c)
}
