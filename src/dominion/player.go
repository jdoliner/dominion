package dominion

import (
	"fmt"
	"sort"

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

func (p *Player) Cleanup() {
	p.Hand.MoveTo(p.Discard)
	p.Play.MoveTo(p.Discard)
}

// Play p.Hand[card]
func (p *Player) PlayCards(cards ...int) {
	sort.Ints(cards)
	var hand Cards
	for i, c := range p.Hand {
		if len(cards) == 0 {
			break
		}
		if i == cards[0] {
			p.Play = append(p.Play, c)
			cards = cards[1:]
		} else {
			hand = append(hand, c)
		}
	}
	if len(cards) != 0 {
		panic(fmt.Sprintf("%d is out of range of hand (len %d)", cards[0], len(p.Hand)))
	}
	p.Hand = hand
}

func (p *Player) PlayTreasures() {
	var ts []int
	for i, c := range p.Hand {
		if c.Type.Is(Treasure) {
			ts = append(ts, i)
		}
	}
	p.PlayCards(ts...)
}

type Choice struct {
	Make        func(*Game)
	Description string
}

type Decision []*Choice
