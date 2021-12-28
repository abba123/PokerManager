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
	sqldb, _ := db.DB()
	defer sqldb.Close()
	//產生table
	db.Debug().AutoMigrate(&Game{})
	db.Debug().AutoMigrate(&User{})
	db.Migrator()
}

func InsertUserDB(username string, password string) error {
	db := ConnectDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	user := User{Username: username, Password: password}
	err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user).Error
	return err
}

func GetUserDB(username string) User {
	db := ConnectDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	user := User{}
	db.First(&user, "username = ?", username)

	return user
}

func InsertHandDB(name string, tables []poker.Table) {
	db := ConnectDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	games := []Game{}
	for _, table := range tables {
		for _, value := range table.Player {
			game := Game{}
			game.ID = table.ID
			game.Time = table.Time
			db.FirstOrCreate(&game.User, User{Username: name})
			db.FirstOrCreate(&game.Player, Player{Playername: value.Name})
			if len(value.Card) > 0 {
				db.FirstOrCreate(&game.HeroCard1, Card{Num: value.Card[0].Num, Suit: value.Card[0].Suit})
				db.FirstOrCreate(&game.HeroCard2, Card{Num: value.Card[1].Num, Suit: value.Card[1].Suit})
			}
			game.Gain = value.Gain
			if value.Action.Preflop != "" {
				db.FirstOrCreate(&game.Preflop, Action{Action: value.Action.Preflop})
			}
			if value.Action.Flop != "" {
				db.FirstOrCreate(&game.Flop, Action{Action: value.Action.Flop})
			}
			if value.Action.Turn != "" {
				db.FirstOrCreate(&game.Turn, Action{Action: value.Action.Turn})
			}
			if value.Action.River != "" {
				db.FirstOrCreate(&game.River, Action{Action: value.Action.River})
			}

			db.FirstOrCreate(&game.Seat, Seat{Seat: value.Seat})

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
			games = append(games, game)
		}
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&games)
}

func GetGainDB(gain string, username string) []Game {

	games := []Game{}

	db := ConnectDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	db = db.Joins(" INNER JOIN  users ON games.user_id = users.id ").Where("username = ?", username).Order("time")
	db = db.Joins(" INNER JOIN  players ON games.player_id = players.id ").Where("playername = ?", "Hero")
	if gain != "all" {
		g, _ := strconv.ParseFloat(gain[1:], 64)
		db = db.Where("gain >= ?", g)
	}

	db.Preload(clause.Associations).Find(&games)

	return games
}

func GetSeatDB(seat string, username string) []Game {

	games := []Game{}

	db := ConnectDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	db = db.Joins(" INNER JOIN  users ON games.user_id = users.id ").Where("username = ?", username).Order("time")
	db = db.Joins(" INNER JOIN  players ON games.player_id = players.id ").Where("playername = ?", "Hero").Order("time")

	if seat != "all" {
		db = db.Joins("INNER JOIN  seats ON games.seat_id = seats.id").Where("seat = ?", seat)
	}

	db.Preload(clause.Associations).Find(&games)

	return games
}

func GetProfitDB(username string, player string) []float64 {
	var results []float64

	db := ConnectDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	db = db.Table("games").Joins(" INNER JOIN  users ON games.user_id = users.id ").Where("username = ?", username)
	db = db.Joins(" INNER JOIN  players ON games.player_id = players.id ").Where("playername = ?", player)
	db.Select("gain").Order("time").Scan(&results)

	return results
}

func GetActionDB(stage string, action string, username string, player string) float64 {
	var result int64

	db := ConnectDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	db = db.Table("games").Joins(" INNER JOIN  users ON games.user_id = users.id ").Where("username = ?", username)
	db = db.Joins(" INNER JOIN  players ON games.player_id = players.id ").Where("playername = ?", player)
	db = db.Joins(" INNER JOIN  actions ON games."+stage+"_id = actions.id ").Where("action LIKE ?", action+"%")
	db.Count(&result)

	return float64(result)
}

func GetPlayerDB(username string) []string {
	result := []string{}
	db := ConnectDB()
	sqldb, _ := db.DB()
	defer sqldb.Close()

	db = db.Table("games").Joins(" INNER JOIN  users ON games.user_id = users.id ").Where("username = ?", username)
	db = db.Joins(" INNER JOIN  players ON games.player_id = players.id ").Select("players.playername").Distinct("players.playername")
	db.Scan(&result)
	return result
}
