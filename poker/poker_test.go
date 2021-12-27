package poker

import (
	"os"
	"testing"
)

func TestStraightFlush(t *testing.T) {

	player := Player{}
	table := Table{}

	player.Card = []Card{{Num: 1, Suit: "c"}, {Num: 9, Suit: "c"}}
	table.Card = []Card{{Num: 10, Suit: "c"}, {Num: 11, Suit: "c"}, {Num: 12, Suit: "c"}, {Num: 13, Suit: "c"}, {Num: 3, Suit: "c"}}

	rank, _ := GetRank(player, table)
	if rank == 8 {
		t.Log("Straight flush PASS")
	} else {
		t.Error("Straight flush FAIL")
	}
}

func TestFourKind(t *testing.T) {

	player := Player{}
	table := Table{}

	player.Card = []Card{{Num: 1, Suit: "c"}, {Num: 1, Suit: "s"}}
	table.Card = []Card{{Num: 1, Suit: "h"}, {Num: 1, Suit: "d"}, {Num: 12, Suit: "c"}, {Num: 13, Suit: "c"}, {Num: 3, Suit: "c"}}

	rank, _ := GetRank(player, table)

	if rank == 7 {
		t.Log("Four of a kind PASS")
	} else {
		t.Error("Four of a kind FAIL")
	}
}

func TestFullHouse(t *testing.T) {

	player := Player{}
	table := Table{}

	player.Card = []Card{{Num: 1, Suit: "c"}, {Num: 1, Suit: "s"}}
	table.Card = []Card{{Num: 1, Suit: "h"}, {Num: 12, Suit: "d"}, {Num: 12, Suit: "c"}, {Num: 13, Suit: "c"}, {Num: 3, Suit: "c"}}

	rank, _ := GetRank(player, table)

	if rank == 6 {
		t.Log("Full house PASS")
	} else {
		t.Error("Full house FAIL")
	}
}

func TestFlush(t *testing.T) {

	player := Player{}
	table := Table{}

	player.Card = []Card{{Num: 1, Suit: "c"}, {Num: 1, Suit: "s"}}
	table.Card = []Card{{Num: 2, Suit: "c"}, {Num: 12, Suit: "d"}, {Num: 12, Suit: "c"}, {Num: 13, Suit: "c"}, {Num: 3, Suit: "c"}}

	rank, _ := GetRank(player, table)

	if rank == 5 {
		t.Log("Flush PASS")
	} else {
		t.Error("Flush FAIL")
	}
}

func TestStraight(t *testing.T) {

	player := Player{}
	table := Table{}

	player.Card = []Card{{Num: 1, Suit: "c"}, {Num: 9, Suit: "s"}}
	table.Card = []Card{{Num: 11, Suit: "c"}, {Num: 10, Suit: "d"}, {Num: 12, Suit: "c"}, {Num: 13, Suit: "d"}, {Num: 3, Suit: "c"}}

	rank, _ := GetRank(player, table)

	if rank == 4 {
		t.Log("Straight PASS")
	} else {
		t.Error("Straight FAIL", rank)
	}
}

func TestThreeKind(t *testing.T) {

	player := Player{}
	table := Table{}

	player.Card = []Card{{Num: 1, Suit: "c"}, {Num: 1, Suit: "s"}}
	table.Card = []Card{{Num: 1, Suit: "d"}, {Num: 10, Suit: "d"}, {Num: 12, Suit: "c"}, {Num: 11, Suit: "d"}, {Num: 3, Suit: "c"}}

	rank, _ := GetRank(player, table)

	if rank == 3 {
		t.Log("Three of a kind PASS")
	} else {
		t.Error("Three of a kind FAIL")
	}
}

func TestTwoPairs(t *testing.T) {

	player := Player{}
	table := Table{}

	player.Card = []Card{{Num: 1, Suit: "c"}, {Num: 9, Suit: "s"}}
	table.Card = []Card{{Num: 5, Suit: "c"}, {Num: 4, Suit: "h"}, {Num: 10, Suit: "h"}, {Num: 4, Suit: "c"}, {Num: 10, Suit: "s"}}

	rank, _ := GetRank(player, table)

	if rank == 2 {
		t.Log("Two pairs PASS")
	} else {
		t.Error("Two pairs FAIL")
	}
}

func TestPair(t *testing.T) {

	player := Player{}
	table := Table{}

	player.Card = []Card{{Num: 9, Suit: "c"}, {Num: 1, Suit: "s"}}
	table.Card = []Card{{Num: 13, Suit: "c"}, {Num: 6, Suit: "s"}, {Num: 4, Suit: "c"}, {Num: 8, Suit: "s"}, {Num: 8, Suit: "c"}}

	rank, _ := GetRank(player, table)

	if rank == 1 {
		t.Log("Pair PASS")
	} else {
		t.Error("Pair FAIL")
	}
}

func TestHighCard1(t *testing.T) {

	player := Player{}
	table := Table{}

	player.Card = []Card{{Num: 5, Suit: "c"}, {Num: 1, Suit: "s"}}
	table.Card = []Card{{Num: 10, Suit: "s"}, {Num: 4, Suit: "d"}, {Num: 12, Suit: "c"}, {Num: 11, Suit: "d"}, {Num: 3, Suit: "c"}}

	rank, _ := GetRank(player, table)

	if rank == 0 {
		t.Log("High card PASS")
	} else {
		t.Error("High card FAIL")
	}
}

func TestHighCard2(t *testing.T) {

	player := Player{}
	table := Table{}

	player.Card = []Card{{Num: 1, Suit: "c"}, {Num: 5, Suit: "s"}}
	table.Card = []Card{{Num: 10, Suit: "s"}, {Num: 4, Suit: "d"}, {Num: 12, Suit: "c"}, {Num: 11, Suit: "d"}, {Num: 3, Suit: "c"}}

	rank, _ := GetRank(player, table)

	if rank == 0 {
		t.Log("High card PASS")
	} else {
		t.Error("High card FAIL")
	}
}

func TestGetWinRate(t *testing.T) {

	player1 := Player{}
	player2 := Player{}
	player1.Card = []Card{{Num: 1, Suit: "c"}, {Num: 5, Suit: "s"}}
	player2.Card = []Card{{Num: 1, Suit: "h"}, {Num: 4, Suit: "s"}}
	player1.Name = "player1"
	player2.Name = "player2"

	result := GetWinRate([]Player{player1, player2}, 10000)
	if result[player1.Name] > result[player2.Name] {
		t.Log("GetWinRate PASS")
	} else {
		t.Log("GetWinRate FAIL")
	}
}

func TestParseFile(t *testing.T) {
	data, _ := os.ReadFile("testData.txt")
	result := Parsefile("test", string(data))[0]

	table := Table{}
	table.ID = 376515665

	if table.ID == result.ID {
		t.Log("Parse PASS")
	} else {
		t.Log("Parse FAIL")
	}
}
