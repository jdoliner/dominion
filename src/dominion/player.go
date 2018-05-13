package dominion

import (
	"github.com/golang/go/src/math/rand"
)

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

func (p *Player) Shuffle() {
	rand.Shuffle(len(p.Deck), func(i, j int) { p.Deck[i], p.Deck[j] = p.Deck[j], p.Deck[i] })
}

func (p *Player) Reshuffle() {
	p.Deck = append(p.Deck, p.Discard...)
	p.Discard = nil
	p.Shuffle()
}

func (p *Player) Draw(n int) {
	l := len(p.Deck)
	if l < n {
		if l > 0 {
			p.Draw(l)
		}
		if len(p.Discard) > 0 {
			p.Reshuffle()
			p.Draw(n - l)
		}
		return
	}
	p.Hand = append(p.Hand, p.Deck[:n]...)
	p.Deck = p.Deck[n:]
}

type Choice struct {
	Make        func(*Game)
	Description string
}

type Decision []*Choice
