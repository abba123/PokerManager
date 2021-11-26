package main

import (
	"fmt"
	"poker/poker"
)

func main() {

	p1 := poker.Player{}
	p2 := poker.Player{}

	p1.Card = []poker.Card{{Num: 1, Suit: "c"}, {Num: 13, Suit: "c"}}
	p2.Card = []poker.Card{{Num: 5, Suit: "c"}, {Num: 5, Suit: "s"}}

	fmt.Println(poker.GetWinRate(p1, p2, 10000))

}
