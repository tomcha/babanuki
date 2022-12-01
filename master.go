package main

import (
	"fmt"
)

type Master struct {
	players_ []Player
}

func (m *Master) PrepareGame(cards *Hand) {
	fmt.Println("カードを配ります")
	cards.Shuffle()
	numOfCards := cards.GetNumberOfCard()
	numOfPlayers := len(m.players_)
	for i := 0; i < numOfCards; i++ {
		m.players_[i%numOfPlayers].DealCrad(cards.PickCard())
	}
}

func (m *Master) DeclareWin(p Player) {
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

func (m *Master) RegisterPlayer(p Player) {
	m.players_ = append(m.players_, p)
}

func (m *Master) StartGame() {
	fmt.Println("ババ抜きを開始します")
	playersSize := len(m.players_)
	for i := 0; playersSize > 1; i++ {
		playersSize = len(m.players_)
		playerIndex := i % playersSize
		nextPlayerIndex := (i + 1) % playersSize
		fmt.Printf("%s さんの番です\n", m.players_[playerIndex])
		m.players_[playerIndex].Play(&m.players_[nextPlayerIndex])
		m.LookAllPlayersHand()
		//		time.Sleep(time.Second * 2)
	}
	fmt.Println("ババ抜きを終了します")
}

func (m *Master) LookAllPlayersHand() {
	ps := len(m.players_)
	for i := 0; i < ps; i++ {
		fmt.Printf("%s :", m.players_[i])
		fmt.Printf(" %s\n", m.players_[i].myHand.hand_)
	}
}
