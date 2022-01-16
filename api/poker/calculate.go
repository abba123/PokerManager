package poker

import (
	"math/rand"
	"sort"
	"time"

	"github.com/shopspring/decimal"
)

func GetWinRate(player []Player, times int) map[string]float64 {

	result := map[string]float64{}
	var total int64

	for i := 0; i < times; i++ {
		winner := GetWinner(player, i)
		for _, w := range winner {
			result[w.Name] += 1
			total += 1
		}
	}

	for k, v := range result {
		result[k], _ = decimal.NewFromFloat(v).Div(decimal.NewFromInt(total)).Float64()
	}

	return result
}

func GetWinner(player []Player, secNum int) []Player {
	cardSet := initCardSet()

	for _, p := range player {
		cardSet = removeCard(cardSet, p.Card[0])
		cardSet = removeCard(cardSet, p.Card[1])
	}

	table := Table{}
	for i := 0; i < 5; i++ {
		card := Card{}
		card, cardSet = generateCard(cardSet, secNum+i)
		table.Card = append(table.Card, card)
	}

	for i := range player {
		player[i].Rank, player[i].RankValue = GetRank(player[i], table)
	}

	sort.SliceStable(player, func(i, j int) bool {
		return Bigger(player[i].RankValue, player[j].RankValue)
	})

	sort.SliceStable(player, func(i, j int) bool {
		return player[i].Rank > player[j].Rank
	})

	newPlayer := []Player{}

	for i := 1; i < len(player); i++ {
		if player[i].Rank == player[i-1].Rank && Same(player[i].RankValue, player[i-1].RankValue) {
			continue
		} else {
			newPlayer = player[:i]
			break
		}
	}

	return newPlayer
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
