package model

import (
	"time"
)

type User struct {
	ID       int    `gorm:"type:int;primaryKey;" json:"ID,omitempty"`
	Username string `gorm:"type:varchar(100);unique" json:"username,omitempty"`
	Password string `gorm:"type:varchar(100)" json:"password,omitempty"`
}

type Game struct {
	//gorm為model的tag標籤，v2版的auto_increment要放在type裡面，v1版是放獨立定義
	ID         int       `gorm:"type:int;primaryKey;autoIncrement:false" json:"ID,omitempty"`
	PlayerID   int       `gorm:"default:NULL;"`
	Player     User      `gorm:"foreignKey:PlayerID;association:ID;"`
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
