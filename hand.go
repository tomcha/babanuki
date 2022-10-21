package main

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
