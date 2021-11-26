package main

import (
	//"fmt"
	"poker/poker"
)

func main() {

	t := poker.Table{}

	p1 := poker.Player{Name: "1"}
	p2 := poker.Player{Name: "2"}
	p3 := poker.Player{Name: "3"}

	p1.Card = []poker.Card{{Num: 1, Suit: "c"}, {Num: 13, Suit: "c"}}
	p2.Card = []poker.Card{{Num: 5, Suit: "c"}, {Num: 6, Suit: "s"}}
	p3.Card = []poker.Card{{Num: 4, Suit: "c"}, {Num: 3, Suit: "c"}}

	t.Player = append(t.Player, p2)
	t.Player = append(t.Player, p1)
	t.Player = append(t.Player, p3)

	poker.GetWinRate(t.Player, 1000)

}
