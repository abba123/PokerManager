package model

import (
	"poker/poker"
	"testing"
	"time"
)

func TestConnectDB(t *testing.T) {
	//viper.Set("DATABASE", "127.0.0.1")
	db := ConnectDB("testdb")
	if db != nil {
		t.Log("Connect DB PASS")
	} else {
		t.Error("Connect DB FAIL")
	}
}

func TestInitDB(t *testing.T) {
	err := InitDB("testdb")

	if err == nil {
		t.Log("Init DB PASS")
	} else {
		t.Log("Init DB FAIL")
	}
}

func TestUserDB(t *testing.T) {
	InsertUserDB("testdb", "test", "test")
	user := GetUserDB("testdb", "test")

	if user.Password == "test" {
		t.Log("User DB PASS")
	} else {
		t.Error("User DB FAIL")
	}
}

func InsertTestHandDB() {
	player := poker.Player{
		Name: "Hero",
		Seat: "BB",
		Gain: 1.0,
		Action: struct {
			Preflop string
			Flop    string
			Turn    string
			River   string
		}{Preflop: "X", Flop: "X", Turn: "X", River: "X"},
		Card: []poker.Card{{Num: 1, Suit: "c"}, {Num: 9, Suit: "c"}},
	}
	table := poker.Table{
		ID:     1,
		Time:   time.Now(),
		Player: map[string]poker.Player{"test": player},
		Card:   []poker.Card{{Num: 10, Suit: "c"}, {Num: 11, Suit: "c"}, {Num: 12, Suit: "c"}, {Num: 13, Suit: "c"}, {Num: 3, Suit: "c"}},
	}

	InsertHandDB("testdb", "test", []poker.Table{table})
}
func TestGetPlayerDB(t *testing.T) {
	InsertTestHandDB()
	result := GetPlayerDB("testdb", "test")
	if len(result) != 0 {
		t.Log("InsertHand DB PASS")
	} else {
		t.Error("InsertHand DB FAIL")
	}
}

func TestGetGainDB(t *testing.T) {
	InsertTestHandDB()
	result := GetGainDB("testdb", "all", "test")

	if result[0].Gain == 1.0 {
		t.Log("Get Gain PASS")
	} else {
		t.Error("Get Gain FAIL")
	}
}

func TestGetSeatDB(t *testing.T) {
	InsertTestHandDB()
	result := GetSeatDB("testdb", "all", "test")
	if result[0].Seat.Seat == "BB" {
		t.Log("Get Seat PASS")
	} else {
		t.Error("Get Seat FAIL")
	}
}

func TestGetProfitDB(t *testing.T) {
	InsertTestHandDB()
	result := GetProfitDB("testdb", "test", "Hero")

	if result[0] == 1.0 {
		t.Log("Get Profit PASS")
	} else {
		t.Error("Get Profit FAIL")
	}
}

func TestGetActionDB(t *testing.T) {
	InsertTestHandDB()
	result := GetActionDB("testdb", "Flop", "X", "test", "Hero")

	if result == 1.0 {
		t.Log("Get Action PASS")
	} else {
		t.Error("Get Action FAIL")
	}
}
