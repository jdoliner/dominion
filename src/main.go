package main

import (
	"fmt"
	d "github.com/jdoliner/dominion/src/dominion"
)

func main() {
	g := d.NewGame(2)
	fmt.Printf("%+v\n", g)
}
