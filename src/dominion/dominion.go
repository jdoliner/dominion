package dominion

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Game struct {
	Players   []*Player
	Trash     []*Card
	Decisions []*Decision
}

func NewGame(players int) *Game {
	result := &Game{}
	for i := 0; i < players; i++ {
		result.Players = append(result.Players, NewPlayer(StartDeck))
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

type Player struct {
	Actions int
	Coin    int
	Buys    int
	Deck    Cards
	Hand    Cards
	Discard Cards
	Play    Cards
}

func NewPlayer(deck []*Card) *Player {
	return &Player{Deck: deck}
}

func (p *Player) String() string {
	var result string
	result += "Deck: " + p.Deck.String() + "\n"
	result += "Hand: " + p.Hand.String() + "\n"
	result += "Play: " + p.Play.String() + "\n"
	return result
}

type Type int64

const (
	Victory Type = 1 << iota
	Treasure
	Action
	Curse
)

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

type Choice struct {
	Make        func(*Game)
	Description string
}

type Decision []*Choice
