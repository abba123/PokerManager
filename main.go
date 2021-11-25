package main

import (
	"fmt"
	"poker/poker"
)

func main() {
	p := poker.Player{}
	t := poker.Table{}

	p.Card = []poker.Card{{Num: 1, Suit: "c"}, {Num: 9, Suit: "s"}}
	t.Card = []poker.Card{{Num: 10, Suit: "c"}, {Num: 11, Suit: "c"}, {Num: 12, Suit: "c"}, {Num: 13, Suit: "c"}, {Num: 3, Suit: "c"}}

	rank, _ := poker.GetRank(p, t)
	fmt.Println(rank)
}
