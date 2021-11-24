package poker

/*
{
	Straight flush		800
	Four of a kind		700
	Full house			600
	Flush				500
	Straight			400
	Three of a kind		300
	Two pairs			200
	Pair				100
	High card			0
}
*/

func GetRank(player Player, table Table) int {

	var card []Card
	//var maxSuit string
	var maxSuitCount int
	var maxValue int

	card = append(card, player.Card...)
	card = append(card, table.Card...)

	maxSuitCount, maxValue = getSuits(card)

	if maxSuitCount == 5 {
		return 500 + maxValue
	} else {
		return 0
	}
}

func getSuits(card []Card) (int, int) {
	suits := map[string]int{}
	var maxSuitCount int
	var maxValue int

	for _, c := range card {
		suits[c.Suit] += 1
		if maxSuitCount < suits[c.Suit] {
			maxSuitCount = suits[c.Suit]
			if maxValue < c.Num {
				maxValue = c.Num
			}
		}
	}

	return maxSuitCount, maxValue
}
