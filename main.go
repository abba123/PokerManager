package main

import (
	//"fmt"
	"poker/poker"
)

func main() {
	p := poker.Player{}
	t := poker.Table{}

	p.Card = append(p.Card, poker.Card{Num: 1,Suit: "s"})
	p.Card = append(p.Card, poker.Card{Num: 2,Suit: "h"})

	t.Card = append(t.Card, poker.Card{Num: 2,Suit: "s"})
	t.Card = append(t.Card, poker.Card{Num: 2,Suit: "h"})
	t.Card = append(t.Card, poker.Card{Num: 2,Suit: "d"})
	t.Card = append(t.Card, poker.Card{Num: 2,Suit: "c"})
	t.Card = append(t.Card, poker.Card{Num: 2,Suit: "c"})

	poker.GetRank(p,t)

}
