package dominion

var (
	Estate   = &Card{Name: "Estate", Type: Victory, Cost: 0, Score: 1}
	Duchy    = &Card{Name: "Duchy", Type: Victory, Cost: 5, Score: 3}
	Province = &Card{Name: "Provence", Type: Victory, Cost: 8, Score: 6}

	Copper = &Card{Name: "Copper", Type: Treasure, Cost: 0, Score: 0}
	Silver = &Card{Name: "Silver", Type: Treasure, Cost: 3, Score: 0}
	Gold   = &Card{Name: "Gold", Type: Treasure, Cost: 6, Score: 0}

	StartDeck = Cards{
		Estate, Estate, Estate,
		Copper, Copper, Copper, Copper, Copper, Copper, Copper,
	}
	StartKingdom = []Cards{
		NewCards(Estate, 8), NewCards(Duchy, 8), NewCards(Province, 8),
		NewCards(Copper, 46), NewCards(Silver, 40), NewCards(Gold, 30),
	}
)
