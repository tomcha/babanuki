package main

type Player struct {
	name_  string
	myHand *Hand
	table_ *Table
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

func (p *Player) ShowHand() Hand {
	if len(p.myHand.hand_) == 0 {
		//master
	}
	p.myHand.Shuffle()
	return *p.myHand
}
