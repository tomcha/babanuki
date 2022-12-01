package main

import "fmt"

type Player struct {
	name_   string
	myHand  *Hand
	table_  *Table
	master_ *Master
}

func (p Player) String() string {
	return p.name_
}

func (p *Player) RecieveCard(c Card) {
	p.myHand.AddCard(c)
	sc := p.myHand.FindSameNumberCard()
	if len(sc) != 0 {
		p.table_.DisposeCard(sc[0])
		p.table_.DisposeCard(sc[1])
	}
}

func (p *Player) ShowHand() *Hand {
	if len(p.myHand.hand_) == 1 {
		p.master_.DeclareWin(*p)
		return p.myHand
	}
	p.myHand.Shuffle()
	return p.myHand
}

func (p *Player) DealCrad(c Card) {
	p.myHand.hand_ = append(p.myHand.hand_, c)
	sameCards := p.myHand.FindSameNumberCard()
	if len(sameCards) != 0 {
		p.table_.DisposeCard(sameCards[0])
		p.table_.DisposeCard(sameCards[1])
	}
}

func (p *Player) Play(nextPalyer *Player) {
	nextHand := nextPalyer.ShowHand()
	pickedCard := nextHand.PickCard()
	fmt.Printf("%s から %s のカードを引きました\n", nextPalyer.String(), pickedCard.String())
	p.DealCrad(pickedCard)
	if p.myHand.GetNumberOfCard() == 0 {
		p.master_.DeclareWin(*p)
	} else {
		fmt.Printf("%s : 残りの手札は %s です\n", p.String(), p.myHand.hand_)
	}
}
