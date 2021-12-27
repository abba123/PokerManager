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

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("abba123:abbaABBA123@tcp("+mysqlIP+")/pokerdb?parseTime=true"), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return db
	}

	return db
}

func InitDB() {

	//連接MySQL
	db := ConnectDB()

	//產生table
	db.Debug().AutoMigrate(&Game{})
	db.Debug().AutoMigrate(&User{})
	db.Migrator()
}

func InsertUserDB(username string, password string) error {
	db := ConnectDB()

	user := User{Username: username, Password: password}
	err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user).Error
	return err
}

func GetUserDB(username string) User {
	db := ConnectDB()

	user := User{}
	db.First(&user, "username = ?", username)

	return user
}

func InsertHandDB(tables []poker.Table) {
	db := ConnectDB()

	games := []Game{}
	for _, table := range tables {
		game := Game{}
		game.ID = table.ID
		game.Time = table.Time
		db.FirstOrCreate(&game.Player, User{Username: table.Player[0].Name})
		db.FirstOrCreate(&game.HeroCard1, Card{Num: table.Player[0].Card[0].Num, Suit: table.Player[0].Card[0].Suit})
		db.FirstOrCreate(&game.HeroCard2, Card{Num: table.Player[0].Card[1].Num, Suit: table.Player[0].Card[1].Suit})

		if len(table.Card) > 0 {
			db.FirstOrCreate(&game.TableCard1, Card{Num: table.Card[0].Num, Suit: table.Card[0].Suit})
		}
		if len(table.Card) > 1 {
			db.FirstOrCreate(&game.TableCard2, Card{Num: table.Card[1].Num, Suit: table.Card[1].Suit})
		}
		if len(table.Card) > 2 {
			db.FirstOrCreate(&game.TableCard3, Card{Num: table.Card[2].Num, Suit: table.Card[2].Suit})
		}
		if len(table.Card) > 3 {
			db.FirstOrCreate(&game.TableCard4, Card{Num: table.Card[3].Num, Suit: table.Card[3].Suit})
		}
		if len(table.Card) > 4 {
			db.FirstOrCreate(&game.TableCard5, Card{Num: table.Card[4].Num, Suit: table.Card[4].Suit})
		}
		db.FirstOrCreate(&game.Seat, Seat{Seat: table.Player[0].Seat})

		game.Gain = table.Player[0].Gain
		if table.Player[0].Action.Preflop != "" {
			db.FirstOrCreate(&game.Preflop, Action{Action: table.Player[0].Action.Preflop})
		}
		if table.Player[0].Action.Flop != "" {
			db.FirstOrCreate(&game.Flop, Action{Action: table.Player[0].Action.Flop})
		}
		if table.Player[0].Action.Turn != "" {
			db.FirstOrCreate(&game.Turn, Action{Action: table.Player[0].Action.Turn})
		}
		if table.Player[0].Action.River != "" {
			db.FirstOrCreate(&game.River, Action{Action: table.Player[0].Action.River})
		}

		games = append(games, game)
	}
	db.Create(&games)
}

func GetGainDB(gain string, player string) []Game {

	games := []Game{}

	db := ConnectDB()
	db = db.Joins(" INNER JOIN  users ON games.player_id = users.id ").Where("username = ?", player).Order("time")
	if gain != "all" {
		g, _ := strconv.ParseFloat(gain[1:], 64)
		db.Where("gain >= ?", g)
	}
	
	db.Preload(clause.Associations).Find(&games)

	return games
}

func GetSeatDB(seat string, player string) []Game {

	games := []Game{}

	db := ConnectDB()
	db = db.Debug().Joins("INNER JOIN  users ON games.player_id = users.id").Where("username = ?", player).Order("time")

	if seat != "all" {
		db.Joins("INNER JOIN  seats ON games.seat_id = seats.id").Where("seat = ?", seat)
	}

	db.Preload(clause.Associations).Find(&games)
	return games
}

func GetProfitDB(player string) []float64 {
	var results []float64

	db := ConnectDB()

	db.Table("games").Joins(" INNER JOIN  users ON games.player_id = users.id ").Where("username = ?", player).Select("gain").Order("time").Scan(&results)

	return results
}

func GetActionDB(stage string, action string, player string) float64 {
	var result int64

	db := ConnectDB()

	db.Table("games").Joins("INNER JOIN  users ON games.player_id = users.id").Joins(" INNER JOIN  actions ON games."+stage+"_id = actions.id ").Where("username = ?", player).Where("action LIKE ?", action+"%").Count(&result)

	return float64(result)
}
