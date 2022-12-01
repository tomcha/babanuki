package main

func main() {
	master := Master{
		players_: []Player{},
	}
	table := Table{
		cards_: []Card{},
	}
	murata := Player{
		name_:   "murata",
		myHand:  &Hand{},
		table_:  &table,
		master_: &master,
	}
	saito := Player{
		name_:   "saito",
		myHand:  &Hand{},
		table_:  &table,
		master_: &master,
	}
	yamada := Player{
		name_:   "yamada",
		myHand:  &Hand{},
		table_:  &table,
		master_: &master,
	}

	master.RegisterPlayer(murata)
	master.RegisterPlayer(saito)
	master.RegisterPlayer(yamada)

	gameCards := CreateTrump()
	master.PrepareGame(&gameCards)
	master.StartGame()
}

func CreateTrump() (cards Hand) {
	cards.hand_ = []Card{
		{SUIT_CLUB, 1},
		{SUIT_CLUB, 2},
		{SUIT_CLUB, 3},
		{SUIT_CLUB, 4},
		{SUIT_CLUB, 5},
		{SUIT_CLUB, 6},
		{SUIT_CLUB, 7},
		{SUIT_CLUB, 8},
		{SUIT_CLUB, 9},
		{SUIT_CLUB, 10},
		{SUIT_CLUB, 11},
		{SUIT_CLUB, 12},
		{SUIT_CLUB, 13},
		{SUIT_HEART, 1},
		{SUIT_HEART, 2},
		{SUIT_HEART, 3},
		{SUIT_HEART, 4},
		{SUIT_HEART, 5},
		{SUIT_HEART, 6},
		{SUIT_HEART, 7},
		{SUIT_HEART, 8},
		{SUIT_HEART, 9},
		{SUIT_HEART, 10},
		{SUIT_HEART, 11},
		{SUIT_HEART, 12},
		{SUIT_HEART, 13},
		{SUIT_SPADE, 1},
		{SUIT_SPADE, 2},
		{SUIT_SPADE, 3},
		{SUIT_SPADE, 4},
		{SUIT_SPADE, 5},
		{SUIT_SPADE, 6},
		{SUIT_SPADE, 7},
		{SUIT_SPADE, 8},
		{SUIT_SPADE, 9},
		{SUIT_SPADE, 10},
		{SUIT_SPADE, 11},
		{SUIT_SPADE, 12},
		{SUIT_SPADE, 13},
		{SUIT_DIAMOND, 1},
		{SUIT_DIAMOND, 2},
		{SUIT_DIAMOND, 3},
		{SUIT_DIAMOND, 4},
		{SUIT_DIAMOND, 5},
		{SUIT_DIAMOND, 6},
		{SUIT_DIAMOND, 7},
		{SUIT_DIAMOND, 8},
		{SUIT_DIAMOND, 9},
		{SUIT_DIAMOND, 10},
		{SUIT_DIAMOND, 11},
		{SUIT_DIAMOND, 12},
		{SUIT_DIAMOND, 13},
		{JOKER, 0},
	}
	return cards
}
