package poker

import (
	//"fmt"
	"sort"
)

/*
{
	Straight flush		8
	Four of a kind		7
	Full house			6
	Flush				5
	Straight			4
	Three of a kind		3
	Two pairs			2
	Pair				1
	High card			0
}
*/

func GetRank(player Player, table Table) (int, []Card) {

	var card []Card
	var straightValue []Card
	var pairValue []Card
	var maxPairCount int
	var secondPairCount int
	var suitCard []Card

	card = append(card, player.Card...)
	card = append(card, table.Card...)

	for i, c := range card {
		if c.Num == 1 {
			card[i].Num = 14
		}
	}

	sort.SliceStable(card, func(i, j int) bool { return card[i].Num > card[j].Num })

	suitCard = getSuits(card)
	maxPairCount, secondPairCount, pairValue = getPair(card)
	straightValue = ifStraight(card)
	suitCard = getSuits(card)
	if len(suitCard) >= 5 {
		suitStraight := ifStraight(suitCard)
		if(len(suitStraight) >= 5){
			return 8,suitStraight
		}else{
			return 5,suitCard[:5]
		}
	} else if len(straightValue) == 5 {
		return 4,straightValue
	} else {
		if maxPairCount == 4 {
			return 7,pairValue
		} else if maxPairCount == 3 {
			if secondPairCount == 2 {
				return 6,pairValue
			} else {

				return 3,pairValue
			}
		} else if maxPairCount == 2 {
			if secondPairCount == 2 {

				return 2,pairValue
			} else {

				return 1,pairValue
			}
		}
	}

	return 0, card[:5]
}

func getSuits(card []Card) []Card {
	suits := map[string]int{}
	var maxSuitCount int
	var maxValue string
	result := []Card{}

	for _, c := range card {
		suits[c.Suit] += 1
		if maxSuitCount < suits[c.Suit] {
			maxSuitCount = suits[c.Suit]
			maxValue = c.Suit
		}
	}

	for i:=len(card)-1; i>=0 ; i--{
		if(card[i].Suit == maxValue){
			result = append(result, card[i])
		}
	}

	return result
}

func getPair(card []Card) (int,int,[]Card) {
	nums := map[int]int{}

	type pair struct {
		num   int
		count int
	}
	pairs := make([]pair, 0)
	for _, c := range card {
		nums[c.Num] += 1
	}

	for k, v := range nums {
		pairs = append(pairs, pair{num: k, count: v})
	}

	sort.SliceStable(pairs, func(i, j int) bool { return pairs[i].num > pairs[j].num })
	sort.SliceStable(pairs, func(i, j int) bool { return pairs[i].count > pairs[j].count })

	result := []Card{}

	for i:=0; i<len(pairs); i++{
		for j := 0 ;j<pairs[i].count; j++{
			result = append(result, Card{Num: pairs[i].num, Suit: "s"})
			if len(result) == 5{
				break
			}
		}
	}

	return pairs[0].count, pairs[1].count, result
}

func ifStraight(card []Card) []Card{

	for _, c := range card {
		if c.Num == 14 {
			card = append(card, Card{Num: 1, Suit: c.Suit})
		}
	}

	value := []Card{}
	value = append(value, card[0])
	count := 1
	for i := 1; i < len(card); i++ {
		
		if card[i].Num == card[i-1].Num {
			continue
		} else if card[i].Num == card[i-1].Num-1 {
			count += 1
			value = append(value, card[i])
		} else {
			value = append([]Card{}, card[i])
			count = 1
		}

		if count == 5 {
			return value
		}
	}
	return []Card{}
}

func Same(card1 []Card, card2 []Card)bool{
	for i:=0; i<5 ; i++ {
		if card1[i].Num != card2[i].Num{
			return false
		}
	}
	return true
}

func Bigger(card1 []Card, card2 []Card)bool{
	for i:=0; i<5 ; i++ {
		if card1[i].Num > card2[i].Num{
			return true
		}
	}
	return false
}