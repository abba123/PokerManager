package poker

import (
	"fmt"
	"sort"
)

/*
{
	Straight flush		80000
	Four of a kind		70000
	Full house			60000
	Flush				50000
	Straight			40000
	Three of a kind		30000
	Two pairs			20000
	Pair				10000
	High card			0
}
*/

func GetRank(player Player, table Table) (string, int) {

	var card []Card
	var straight bool
	var straightValue []Card
	var maxPairCount int
	var maxPairValue int
	var secondPairCount int
	var secondPairValue int
	var thirdPairValue int
	var suitCount int
	var suitValue int
	var straightSuitCount int
	var straightSuitValue int

	card = append(card, player.Card...)
	card = append(card, table.Card...)
	sort.Slice(card, func(i, j int) bool { return card[i].Num < card[j].Num })
	straightCard := card
	for i, c := range card {
		if c.Num == 1 {
			card[i].Num = 14
			straightCard = append(straightCard, Card{Num: 14, Suit: c.Suit})
		}
	}

	suitCount, suitValue = getSuits(card)
	maxPairValue, maxPairCount, secondPairValue, secondPairCount, thirdPairValue, _ = getPair(card)
	straight, straightValue = ifStraight(straightCard)
	straightSuitCount, straightSuitValue = getSuits(straightValue)
	if straight && straightSuitCount >= 5 {
		return "Straight flush", 8000 + straightSuitValue
	} else if suitCount >= 5 {

		return "Flush", 50000 + suitValue*100
	} else if straight {

		return "Straight", 40000 + straightValue[len(straightValue)-1].Num
	} else {
		if maxPairCount == 4 {

			return "Four of a kind", 70000
		} else if maxPairCount == 3 {
			if secondPairCount == 2 {

				return "Full house", 60000 + maxPairValue*100 + secondPairValue
			} else {

				return "Three of a kind", 30000 + maxPairValue*100 + secondPairValue
			}
		} else if maxPairCount == 2 {
			if secondPairCount == 2 {

				return "Two pairs", 20000 + maxPairValue*100+ secondPairValue*10 + thirdPairValue
			} else {

				return "Pair", 10000 + maxPairValue*100 + secondPairValue
			}
		}
	}

	return "High card", getHandRank(player)
}

func getSuits(card []Card) (int, int) {
	suits := map[string]int{}
	var maxSuitCount int
	var maxValue int

	for _, c := range card {
		suits[c.Suit] += 1
		if maxSuitCount < suits[c.Suit] {
			maxSuitCount = suits[c.Suit]
		}
	}
	count := 0

	fmt.Println(card)
	for i:=len(card)-1; i>=0; i-- {
		if maxSuitCount == suits[card[i].Suit] && count < maxSuitCount{
			maxValue += card[i].Num
			maxValue *= 10
			count += 1
		}
	}

	return maxSuitCount, maxValue
}

func getPair(card []Card) (int, int, int, int, int, int) {
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

	if len(pairs) < 3{
		return pairs[0].num, pairs[0].count, pairs[1].num, pairs[1].count, 0, 0
	}else{
		return pairs[0].num, pairs[0].count, pairs[1].num, pairs[1].count, pairs[2].num, pairs[2].count
	}
}

func ifStraight(card []Card) (bool, []Card) {
	value := make([]Card, 0)
	tmp := value
	result := false
	count := 1
	for i := 1; i < len(card); i++ {
		if card[i].Num == card[i-1].Num {
			tmp = append(tmp, card[i])
			continue
		} else if card[i].Num == card[i-1].Num+1 {
			count += 1
			tmp = append(tmp, card[i])
		} else {
			tmp = nil
			tmp = append(tmp, card[i])
			count = 1
		}

		if count > 5 {
			count -= 1
			tmp = tmp[1:]
		}
		if count == 5 {
			result = true
			value = tmp
		}

	}

	return result, value
}

func getHandRank(player Player) int {
	card := player.Card

	if card[0].Num > card[1].Num {
		return card[0].Num*100 + card[1].Num
	} else {
		return card[1].Num*100 + card[0].Num
	}
}
