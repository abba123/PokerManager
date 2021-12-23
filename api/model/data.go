package model

import (
	"time"
)

type User struct {
	Username string `gorm:"type:varchar(100)" json:"username,omitempty"`
	Password string `gorm:"type:varchar(100)" json:"password,omitempty"`
}

type Seat struct {
	Location string `gorm:"type:varchar(3)" json:"seat,omitempty"`
}

type Card struct {
	Num  int `gorm:"type:int" json:"num,omitempty"`
	Suit string `gorm:"type:varchar(1)" json:"suit,omitempty"`
}

type Action struct {
	Action []string
}

type Game struct {
	//gorm為model的tag標籤，v2版的auto_increment要放在type裡面，v1版是放獨立定義
	ID        int `gorm:"type:int;primaryKey;autoIncrement:false" json:"ID,omitempty"`
	User      User
	Time      time.Time `gorm:"type:TIME" json:"time,omitempty"`
	Seat      Seat
	HeroCard  [2]Card
	TableCard [5]Card
	Gain      float64 `gorm:"type:float" json:"gain,omitempty"`
	Preflop   Action  `gorm:"type:varchar(100)" json:"preflop,omitempty"`
	Flop      Action  `gorm:"type:varchar(100)" json:"flop,omitempty"`
	Turn      Action  `gorm:"type:varchar(100)" json:"turn,omitempty"`
	River     Action  `gorm:"type:varchar(100)" json:"river,omitempty"`
}
