package model

import (
	"poker/api/poker"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func TestConnectDB(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	db := ConnectDB()
	if db != nil {
		t.Log("Connect DB PASS")
	} else {
		t.Error("Connect DB FAIL")
	}
}

func TestInitDB(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	err := InitDB()

	if err == nil {
		t.Log("Init DB PASS")
	} else {
		t.Log("Init DB FAIL")
	}
}

func TestUserDB(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	InsertUserDB("test", "test")
	user := GetUserDB("test")

	if user.Password == "test" {
		t.Log("User DB PASS")
	} else {
		t.Error("User DB FAIL")
	}
}

func InsertTestHandDB() {
	viper.Set("DBNAME", "testdb")
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

	InsertHandDB("test", []poker.Table{table})
}
func TestGetPlayerDB(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	InsertTestHandDB()
	result := GetPlayerDB("test")
	if len(result) != 0 {
		t.Log("InsertHand DB PASS")
	} else {
		t.Error("InsertHand DB FAIL")
	}
}

func TestGetGainDB(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	InsertTestHandDB()
	result := GetGainDB("all", "test")

	if result[0].Gain == 1.0 {
		t.Log("Get Gain PASS")
	} else {
		t.Error("Get Gain FAIL")
	}
}

func TestGetSeatDB(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	InsertTestHandDB()
	result := GetSeatDB("all", "test")
	if result[0].Seat.Seat == "BB" {
		t.Log("Get Seat PASS")
	} else {
		t.Error("Get Seat FAIL")
	}
}

func TestGetProfitDB(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	InsertTestHandDB()
	result := GetProfitDB("test", "Hero")

	if result[0] == 1.0 {
		t.Log("Get Profit PASS")
	} else {
		t.Error("Get Profit FAIL")
	}
}

func TestGetActionDB(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	InsertTestHandDB()
	result := GetActionDB("Flop", "X", "test", "Hero")

	if result == 1.0 {
		t.Log("Get Action PASS")
	} else {
		t.Error("Get Action FAIL")
	}
}

func TestPlayerRedis(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	player := GetPlayerRedis("test")

	if len(player) != 0 {
		t.Log("Insert player redis PASS")
	} else {
		t.Error("Insert player redis FAIL")
	}
	RemoveKeyRedis("test")
}

func TestHandRedis(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	result := GetHandRedis("10", "1.0", "BB", "test")
	if len(result) != 0 {
		t.Log("Hand Redis PASS")
	} else {
		t.Error("Hand Redis FAIL")
	}
	RemoveKeyRedis("test")
}

func TestProfitRedis(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	result := GetProfitRedis("test", "Hero")
	if len(result) != 0 {
		t.Log("Profit Redis PASS")
	} else {
		t.Error("Profit Redis FAIL")
	}
	RemoveKeyRedis("test")
}

func TestActionRedis(t *testing.T) {
	viper.Set("DBNAME", "testdb")
	result := GetActionRedis("Flop", "X", "test", "test")
	if len(result) != 0 {
		t.Log("Action Redis PASS")
	} else {
		t.Error("Action Redis FAIL")
	}
	RemoveKeyRedis("test")
}
