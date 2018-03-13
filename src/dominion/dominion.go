package dominion

type Game struct {
	Players   []*Player
	Trash     []*Card
	Decisions []*Decision
}

func NewGame(players int) *Game {
	result := &Game{}
	for i := 0; i < players; i++ {
		result.Players = append(result.Players, NewPlayer(nil))
	}
	return result
}

type Player struct {
	Actions int
	Coin    int
	Buys    int
	Deck    []*Card
	Hand    []*Card
	Discard []*Card
	Play    []*Card
}

func NewPlayer(deck []*Card) *Player {
	return &Player{Deck: deck}
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

type Choice struct {
	Make        func(*Game)
	Description string
}

type Decision []*Choice
