package main

import (
	"fmt"
	"poker/poker"
)

func main() {
	p := poker.Player{}
	t := poker.Table{}

	p.Card = append(p.Card, poker.Card{Num: 1, Suit: "c"})
	p.Card = append(p.Card, poker.Card{Num: 9, Suit: "s"})

	t.Card = append(t.Card, poker.Card{Num: 10, Suit: "c"})
	t.Card = append(t.Card, poker.Card{Num: 11, Suit: "c"})
	t.Card = append(t.Card, poker.Card{Num: 12, Suit: "c"})
	t.Card = append(t.Card, poker.Card{Num: 13, Suit: "c"})
	t.Card = append(t.Card, poker.Card{Num: 3, Suit: "c"})

	rank, _ := poker.GetRank(p, t)
	fmt.Println(rank)
}
