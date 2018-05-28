package dominion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func countCards(cs Cards) map[string]int {
	result := make(map[string]int)
	for _, c := range cs {
		result[c.Name]++
	}
	return result
}

func TestGame(t *testing.T) {
	g := NewGame(2)
	assert.Equal(t, 2, len(g.Players))
	assert.Equal(t, 0, g.active)
	assert.Equal(t, 0, g.Turn)
	p1Start, p2Start := countCards(append(g.Players[0].Deck, g.Players[0].Hand...)), countCards(append(g.Players[1].Deck, g.Players[1].Hand...))
	assert.Equal(t, 3, p1Start["Estate"])
	assert.Equal(t, 7, p1Start["Copper"])
	assert.Equal(t, 3, p2Start["Estate"])
	assert.Equal(t, 7, p2Start["Copper"])
	assert.Equal(t, 5, len(g.Players[0].Hand))
	assert.Equal(t, 5, len(g.Players[1].Hand))
	assert.Equal(t, 5, len(g.Players[0].Deck))
	assert.Equal(t, 5, len(g.Players[1].Deck))

	g.ActivePlayer().PlayTreasures()
	hand, play := countCards(g.ActivePlayer().Hand), countCards(g.ActivePlayer().Play)
	assert.Equal(t, 0, hand["Copper"])
	assert.Equal(t, 0, play["Estate"])
	assert.Equal(t, 5, hand["Estate"]+play["Copper"])
	g.EndTurn()
	fmt.Printf("%v", g.ActivePlayer())
}
