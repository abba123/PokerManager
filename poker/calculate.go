package poker

import (
	"fmt"
	"math/rand"
	"time"
)

func GetWinRate(p1 Player, p2 Player, times int) float32 {

	win1 := 0
	win2 := 0

	for i := 0; i < times; i++ {
		GetWin(p1, p2, &win1, &win2, i)
	}
	fmt.Println(win1)
	fmt.Println(win2)
	return float32(win1) / float32(times)
}

func GetWin(p1 Player, p2 Player, win1 *int, win2 *int, secNum int) {
	cardSet := initCardSet()

	cardSet = removeCard(cardSet, p1.Card[0])
	cardSet = removeCard(cardSet, p1.Card[1])
	cardSet = removeCard(cardSet, p2.Card[0])
	cardSet = removeCard(cardSet, p2.Card[1])

	table := Table{}
	for i := 0; i < 5; i++ {
		card := Card{}
		card, cardSet = generateCard(cardSet, secNum+i)
		table.Card = append(table.Card, card)
	}

	rank1, value1 := GetRank(p1, table)
	rank2, value2 := GetRank(p2, table)
/*
	fmt.Println(table)
	fmt.Println(p1.Card)
	fmt.Println(p2.Card)
	fmt.Println(rank1, value1)
	fmt.Println(rank2, value2)
*/

	if rank1 > rank2{
		*win1 += 1
	}else if rank2 > rank1{
		*win2 += 1
	}else{
		if !Same(value1, value2){
			if Bigger(value1, value2){
				*win1 += 1
			}else{
				*win2 += 1
			}
		}
	}
}

func initCardSet() []Card {
	cardSet := make([]Card, 0)

	for _, suit := range []string{"s", "h", "d", "c"} {
		for _, num := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13} {
			cardSet = append(cardSet, Card{Suit: suit, Num: num})
		}
	}

	return cardSet
}

func removeCard(cardSet []Card, card Card) []Card {
	newCardSet := []Card{}
	for _, c := range cardSet {
		if c.Suit != card.Suit || c.Num != card.Num {
			newCardSet = append(newCardSet, Card{Suit: c.Suit, Num: c.Num})
		}
	}
	return newCardSet
}

func generateCard(cardSet []Card, secNum int) (Card, []Card) {
	rand.Seed(time.Now().UnixNano() + int64(secNum))
	index := rand.Intn(len(cardSet))

	card := Card{}
	card.Num = (cardSet)[index].Num
	card.Suit = (cardSet)[index].Suit

	cardSet = removeCard(cardSet, card)

	return card, cardSet
}
