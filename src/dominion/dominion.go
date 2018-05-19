package dominion

import (
	"fmt"
	"strings"
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
	result.Kingdom = &Kingdom{Piles: StartKingdom}
	return result
}

func (g *Game) String() string {
	var result string
	for i, p := range g.Players {
		result += fmt.Sprintf("Player %d:\n", i) + p.String() + "\n"
	}
	result += "Kingdom:\n"
	result += g.Kingdom.String() + "\n"
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

type Kingdom struct {
	Piles []Cards
}

func (k *Kingdom) String() string {
	var piles []string
	for _, p := range k.Piles {
		if len(p) > 0 {
			piles = append(piles, fmt.Sprintf("[%s x %d]", p[0].String(), len(p)))
		} else {
			piles = append(piles, "[ ]")
		}
	}
	return strings.Join(piles, "\n")
}
