package main

import (
	"fmt"
	d "github.com/jdoliner/dominion/src/dominion"
)

func main() {
	g := d.NewGame(2)
	fmt.Printf("%+v\n", g)
	g.ActivePlayer().PlayTreasures()
	fmt.Printf("%+v\n", g)
	g.ActivePlayer().Cleanup()
	fmt.Printf("%+v\n", g)
}
