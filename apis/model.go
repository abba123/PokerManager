package api

import (
	"context"
	"encoding/json"
	"fmt"
	"poker/poker"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type user struct {
	Username string `gorm:"type:varchar(100);primaryKey" json:"username,omitempty"`
	Password string `gorm:"type:varchar(100)" json:"password,omitempty"`
}

var ctx = context.Background()

type game struct {
	//gorm為model的tag標籤，v2版的auto_increment要放在type裡面，v1版是放獨立定義
	Player     string    `gorm:"type:varchar(100);primaryKey;autoIncrement:false" json:"player,omitempty"`
	ID         int       `gorm:"type:int;primaryKey;autoIncrement:false" json:"ID,omitempty"`
	Time       time.Time `gorm:"type:TIME" json:"time,omitempty"`
	Seat       string    `gorm:"type:varchar(100)" json:"seat,omitempty"`
	HeroCard1  string    `gorm:"type:varchar(100)" json:"herocard1,omitempty"`
	HeroCard2  string    `gorm:"type:varchar(100)" json:"herocard2,omitempty"`
	TableCard1 string    `gorm:"type:varchar(100)" json:"tablecard1,omitempty"`
	TableCard2 string    `gorm:"type:varchar(100)" json:"tablecard2,omitempty"`
	TableCard3 string    `gorm:"type:varchar(100)" json:"tablecard3,omitempty"`
	TableCard4 string    `gorm:"type:varchar(100)" json:"tablecard4,omitempty"`
	TableCard5 string    `gorm:"type:varchar(100)" json:"tablecard5,omitempty"`
	Gain       float64   `gorm:"type:float" json:"gain,omitempty"`
	Preflop    string    `gorm:"type:varchar(100)" json:"preflop,omitempty"`
	Flop       string    `gorm:"type:varchar(100)" json:"flop,omitempty"`
	Turn       string    `gorm:"type:varchar(100)" json:"turn,omitempty"`
	River      string    `gorm:"type:varchar(100)" json:"river,omitempty"`
}

func InitDB() *gorm.DB {

	//連接MySQL

	//IP := "database-1.crj366caarmq.us-east-2.rds.amazonaws.com"
	IP := "127.0.0.1"

	db, err := gorm.Open(mysql.Open("abba123:abbaABBA123@tcp("+IP+":3306)/pokerdb?parseTime=true"), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return db
	}

	//產生table
	db.Debug().AutoMigrate(&game{})
	db.Debug().AutoMigrate(&user{})
	db.Migrator()

	return db
}

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func InsertUserDB(username string, password string) {
	db := InitDB()

	user := user{Username: username, Password: password}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
}

func GetUserDB(username string) user {
	db := InitDB()

	user := user{}
	db.First(&user, "username = ?", username)

	return user
}

func InsertHandDB(tables []poker.Table) {
	db := InitDB()

	games := []game{}
	for _, table := range tables {
		game := game{}
		game.ID = table.ID
		game.Time = table.Time
		game.Player = table.Player[0].Name
		game.HeroCard1 = strconv.Itoa(table.Player[0].Card[0].Num) + table.Player[0].Card[0].Suit
		game.HeroCard2 = strconv.Itoa(table.Player[0].Card[1].Num) + table.Player[0].Card[1].Suit

		if len(table.Card) > 0 {
			game.TableCard1 = strconv.Itoa(table.Card[0].Num) + table.Card[0].Suit
		}
		if len(table.Card) > 1 {
			game.TableCard2 = strconv.Itoa(table.Card[1].Num) + table.Card[1].Suit
		}
		if len(table.Card) > 2 {
			game.TableCard3 = strconv.Itoa(table.Card[2].Num) + table.Card[2].Suit
		}
		if len(table.Card) > 3 {
			game.TableCard4 = strconv.Itoa(table.Card[3].Num) + table.Card[3].Suit
		}
		if len(table.Card) > 4 {
			game.TableCard5 = strconv.Itoa(table.Card[4].Num) + table.Card[4].Suit
		}

		game.Seat = table.Player[0].Seat
		game.Gain = table.Player[0].Gain
		game.Preflop = strings.Join(table.Player[0].Action.Preflop, " ")
		game.Flop = strings.Join(table.Player[0].Action.Flop, " ")
		game.Turn = strings.Join(table.Player[0].Action.Turn, " ")
		game.River = strings.Join(table.Player[0].Action.River, " ")

		games = append(games, game)

	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&games)

}

func getGainDB(gain string, player string) []game {

	games := []game{}

	db := InitDB()

	db = db.Order("time").Where("player = ?", player)
	if gain != "all" {
		g, _ := strconv.ParseFloat(gain[1:], 64)
		db.Where("gain >= ?", g)
	}
	db.Find(&games)

	return games
}

func getSeatDB(seat string, player string) []game {

	games := []game{}

	db := InitDB()
	db = db.Order("time").Where("player = ?", player)

	if seat != "all" {
		db.Where("seat = ?", seat)
	}

	db.Find(&games)

	return games
}

func getHandRedis(num string, gain string, seat string, player string) []game {
	client := InitRedis()

	existGain, _ := client.Exists(ctx, player+"gain"+gain).Result()
	existSeat, _ := client.Exists(ctx, player+"seat"+seat).Result()

	if existGain == 0 {
		insertGainRedis(gain, player)
	}
	if existSeat == 0 {
		insertSeatRedis(seat, player)
	}
	client.ZInterStore(ctx, "inter", &redis.ZStore{Keys: []string{player + "gain" + gain, player + "seat" + seat}}).Result()
	results, _ := client.ZRange(ctx, "inter", 0, -1).Result()
	client.Del(ctx, "inter")

	games := []game{}
	n, _ := strconv.Atoi(num)
	for i := 0; i < n && i < len(results); i++ {
		result := results[i]
		g := game{}
		json.Unmarshal([]byte(result), &g)
		games = append(games, g)
	}
	return games
}

func insertGainRedis(gain string, player string) {
	games := getGainDB(gain, player)

	client := InitRedis()

	for i, game := range games {
		gameStr, _ := json.Marshal(game)
		client.ZAdd(ctx, player+"gain"+gain, &redis.Z{Score: float64(i), Member: gameStr})
	}
}

func insertSeatRedis(seat string, player string) {
	games := getSeatDB(seat, player)

	client := InitRedis()

	for i, game := range games {
		gameStr, _ := json.Marshal(game)
		client.ZAdd(ctx, player+"seat"+seat, &redis.Z{Score: float64(i), Member: gameStr})
	}
}

func getProfitDB(player string) []float64 {
	var results []float64

	db := InitDB()

	db.Table("games").Where("player = ?", player).Select("gain").Order("time").Scan(&results)

	return results
}

func getProfitRedis(player string) []string {
	client := InitRedis()

	exist, _ := client.Exists(ctx, player+"profit").Result()

	if exist == 0 {
		insertProfitRedis(player)
	}

	result, _ := client.LRange(ctx, player+"profit", 0, -1).Result()

	return result
}

func insertProfitRedis(player string) {
	client := InitRedis()

	profits := getProfitDB(player)

	for _, profit := range profits {
		client.RPush(ctx, player+"profit", fmt.Sprint(profit))
	}

	getProfitRedis(player)
}
