package poker

import (
	"fmt"
	"math/rand"
	"time"
)

func GetWinRate(p1 Player, p2 Player) float32 {

	win1 := 0
	win2 := 0

	for i := 0; i < 100; i++ {
		GetWin(p1, p2, &win1, &win2, i)
	}
	fmt.Println(win1)
	fmt.Println(win2)
	return float32(win1) / float32(100)
}

func GetWin(p1 Player, p2 Player, win1 *int, win2 *int, secNum int) {
	cardSet := initCardSet()

	removeCard(cardSet, p1.Card[0])
	removeCard(cardSet, p1.Card[1])
	removeCard(cardSet, p2.Card[0])
	removeCard(cardSet, p2.Card[1])

	table := Table{}
	for i := 0; i < 5; i++ {
		table.Card = append(table.Card, generateCard(cardSet, secNum))
	}

	rank1, value1 := GetRank(p1, table)
	rank2, value2 := GetRank(p2, table)

	if value1 > value2 {
		*win1 += 1
	} else if value2 > value1 {
		*win2 += 1
	} else {
		fmt.Println(table)
		fmt.Println(p1.Card)
		fmt.Println(p2.Card)
		fmt.Println(rank1, value1)
		fmt.Println(rank2, value2)
	}
}

func initCardSet() map[string][]int {
	cardSet := make(map[string][]int)

	cardSet["s"] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	cardSet["h"] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	cardSet["d"] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	cardSet["c"] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	return cardSet
}

func removeCard(cardSet map[string][]int, card Card) {
	for i, n := range cardSet[card.Suit] {
		if n == card.Num {
			cardSet[card.Suit] = append(cardSet[card.Suit][:i], cardSet[card.Suit][i+1:]...)
			break
		}
	}
}

func generateCard(cardSet map[string][]int, secNum int) Card {
	rand.Seed(time.Now().UnixNano() + int64(secNum))
	index := rand.Intn(len(cardSet["s"]) + len(cardSet["h"]) + len(cardSet["d"]) + len(cardSet["c"]))
	card := Card{}
	if index < len(cardSet["s"]) {
		card.Suit = "s"
		card.Num = cardSet["s"][index]
	} else if index -= len(cardSet["s"]); index < len(cardSet["h"]) {
		card.Suit = "h"
		card.Num = cardSet["h"][index]
	} else if index -= len(cardSet["h"]); index < len(cardSet["d"]) {
		card.Suit = "d"
		card.Num = cardSet["d"][index]
	} else if index -= len(cardSet["d"]); index < len(cardSet["c"]) {
		card.Suit = "c"
		card.Num = cardSet["c"][index]
	}

	removeCard(cardSet, card)

	return card
}
