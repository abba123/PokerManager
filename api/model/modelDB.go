package model

import (
	"fmt"
	"poker/poker"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//const mysqlIP string = "database-1.crj366caarmq.us-east-2.rds.amazonaws.com:3306"
const mysqlIP string = "localhost:3306"


func InitDB() *gorm.DB {

	//連接MySQL

	db, err := gorm.Open(mysql.Open("abba123:abbaABBA123@tcp("+mysqlIP+")/pokerdb?parseTime=true"), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return db
	}

	//產生table
	db.Debug().AutoMigrate(&Game{})
	db.Debug().AutoMigrate(&User{})
	db.Migrator()

	return db
}

func InsertUserDB(username string, password string) {
	db := InitDB()

	user := User{Username: username, Password: password}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
}

func GetUserDB(username string) User {
	db := InitDB()

	user := User{}
	db.First(&user, "username = ?", username)

	return user
}

func InsertHandDB(tables []poker.Table) {
	db := InitDB()

	games := []Game{}
	for _, table := range tables {
		game := Game{}
		game.ID = table.ID
		game.Time = table.Time
		game.User.Username = table.Player[0].Name
		for i := 0; i < 2; i++ {
			game.HeroCard[i].Num = table.Player[0].Card[i].Num
			game.HeroCard[i].Suit = table.Player[0].Card[i].Suit
		}
		for i := 0; i < 5; i++ {
			game.TableCard[i].Num = table.Card[i].Num
			game.TableCard[i].Suit = table.Card[i].Suit
		}
		game.Seat.Location = table.Player[0].Seat
		game.Gain = table.Player[0].Gain
		game.Preflop.action = table.Player[0].Action.Preflop
		game.Flop.action = table.Player[0].Action.Flop
		game.Turn.action = table.Player[0].Action.Turn
		game.River.action = table.Player[0].Action.River

		games = append(games, game)
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&games)
}

func GetGainDB(gain string, player string) []Game {

	games := []Game{}

	db := InitDB()

	db = db.Order("time").Where("player = ?", player)
	if gain != "all" {
		g, _ := strconv.ParseFloat(gain[1:], 64)
		db.Where("gain >= ?", g)
	}
	db.Find(&games)

	return games
}

func GetSeatDB(seat string, player string) []Game {

	games := []Game{}

	db := InitDB()
	db = db.Order("time").Where("player = ?", player)

	if seat != "all" {
		db.Where("seat = ?", seat)
	}

	db.Find(&games)

	return games
}

func GetProfitDB(player string) []float64 {
	var results []float64

	db := InitDB()

	db.Table("games").Where("player = ?", player).Select("gain").Order("time").Scan(&results)

	return results
}

func GetActionDB(stage string, action string, player string) float64 {
	var result int64

	db := InitDB()

	db.Table("games").Where("player = ?", player).Where(stage+" LIKE ?", action+"%").Count(&result)

	return float64(result)
}
