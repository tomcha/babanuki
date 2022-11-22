package main

import (
	"fmt"
)

type Master struct {
	players_ []Player
}

func (m *Master) declareWin(p Player) {
	fmt.Printf("%s is finnished\n", p.name_)
	var ti int
	for i := 0; i < len(m.players_); i++ {
		if m.players_[i] == p {
			ti = i
			break
		}
	}
	if ti == len(m.players_)-1 {
		m.players_ = m.players_[:ti]
	} else {
		m.players_ = append(m.players_[:ti], m.players_[ti+1:]...)
	}
}
