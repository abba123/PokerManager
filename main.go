package main

import (
	"fmt"
	"poker/poker"
)

func main() {

	p1 := poker.Player{}
	p2 := poker.Player{}

	p1.Card = []poker.Card{{Num: 1, Suit: "c"}, {Num: 9, Suit: "s"}}
	p2.Card = []poker.Card{{Num: 2, Suit: "c"}, {Num: 5, Suit: "s"}}
	//poker.GetWinRate(p1, p2)
	fmt.Println(poker.GetWinRate(p1, p2, 1000))

}
