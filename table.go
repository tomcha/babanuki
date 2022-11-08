package main

type Table struct {
	cards_ []Card
}

func (t *Table) disposeCard(c Card) {
	t.cards_ = append(t.cards_, c)
}
