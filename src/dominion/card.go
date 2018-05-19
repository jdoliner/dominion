package dominion

import (
	"strings"

	"github.com/fatih/color"
)

// Card represents a dominion card. It does not represent an actual instance
// of a card in a player's hand, it represents the game data for a card.
type Card struct {
	Name  string
	Type  Type
	Cost  int
	Score int
	Play  func(*Game)
}

func (c *Card) String() string {
	switch c.Type {
	case Victory:
		return color.New(color.FgGreen).SprintFunc()(c.Name)
	case Treasure:
		return color.New(color.FgYellow).SprintFunc()(c.Name)
	case Action:
		return color.New(color.FgWhite).SprintFunc()(c.Name)
	case Curse:
		return color.New(color.FgMagenta).SprintFunc()(c.Name)
	default:
		return c.Name
	}
}

type Cards []*Card

func NewCards(c *Card, n int) Cards {
	var result Cards
	for i := 0; i < n; i++ {
		result = append(result, c)
	}
	return result
}

func (cs Cards) String() string {
	var cards []string
	for _, c := range cs {
		cards = append(cards, c.String())
	}
	return strings.Join(cards, ", ")
}
