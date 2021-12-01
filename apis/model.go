package api

import (
	"fmt"
	"poker/poker"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type user struct {
	Username string `gorm:"type:varchar(100);primaryKey" json:"username,omitempty"`
	Password string `gorm:"type:varchar(100)" json:"password,omitempty"`
}

type game struct {
	//gorm為model的tag標籤，v2版的auto_increment要放在type裡面，v1版是放獨立定義
	ID         int       `gorm:"type:int;primaryKey" json:"ID,omitempty"`
	Time       time.Time `gorm:"type:TIME" json:"time,omitempty"`
	Player     string    `gorm:"type:varchar(100)" json:"player,omitempty"`
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

	db, err := gorm.Open(mysql.Open("abba123:abba123@tcp(127.0.0.1:3306)/pokerdb?parseTime=true"), &gorm.Config{})
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

func SearchHandDB(num int) []game {

	games := []game{}

	db := InitDB()

	db.Limit(num).Find(&games)

	return games
}
