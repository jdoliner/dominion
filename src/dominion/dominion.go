package dominion

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Game struct {
	Players      []*Player
	Trash        Cards
	Kingdom      *Kingdom
	Turn         int
	ActivePlayer int
}

func NewGame(players int) *Game {
	result := &Game{}
	for i := 0; i < players; i++ {
		result.Players = append(result.Players, NewPlayer(StartDeck))
		result.Players[i].Shuffle()
		result.Players[i].Draw(5)
	}
	return result
}

func (g *Game) String() string {
	var result string
	for i, p := range g.Players {
		result += fmt.Sprintf("Player %d:\n", i) + p.String() + "\n"
	}
	return result
}

type Type int64

const (
	Victory Type = 1 << iota
	Treasure
	Action
	Curse
)

func (t Type) Is(u Type) bool {
	return t&u != 0
}

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

func (cs Cards) String() string {
	var cards []string
	for _, c := range cs {
		cards = append(cards, c.String())
	}
	return strings.Join(cards, ", ")
}

type Kingdom struct {
	Piles []Cards
}

func (k *Kingdom) String() string {
	var piles []string
	for _, p := range k.Piles {
		piles = append(piles, p.String())
	}
	return strings.Join(piles, "\n")
}
