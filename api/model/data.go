package model

import (
	"time"
)

type User struct {
	ID       int    `gorm:"type:int;primaryKey;" json:"ID,omitempty"`
	Username string `gorm:"type:varchar(100);unique" json:"username,omitempty"`
	Password string `gorm:"type:varchar(100)" json:"password,omitempty"`
}

type Player struct {
	ID         int    `gorm:"type:int;primaryKey;" json:"ID,omitempty"`
	Playername string `gorm:"type:varchar(100);unique" json:"playername,omitempty"`
}

type Seat struct {
	ID   int    `gorm:"type:int;primaryKey;" json:"ID,omitempty"`
	Seat string `gorm:"type:varchar(3)" json:"seat,omitempty"`
}

type Card struct {
	ID   int    `gorm:"type:int;primaryKey;" json:"ID,omitempty"`
	Num  int    `gorm:"type:int;default:NULL;" json:"num,omitempty"`
	Suit string `gorm:"type:varchar(1);default:NULL;" json:"suit,omitempty"`
}

type Action struct {
	ID     int    `gorm:"type:int;primaryKey;" json:"ID,omitempty"`
	Action string `gorm:"type:varchar(10);default:NULL;" json:"action,omitempty"`
}

type Game struct {
	//gorm為model的tag標籤，v2版的auto_increment要放在type裡面，v1版是放獨立定義
	ID           int       `gorm:"type:int;primaryKey;autoIncrement:false" json:"ID,omitempty"`
	UserID       int       `gorm:"primaryKey"`
	User         User      `gorm:"foreignKey:UserID;association:ID;"`
	PlayerID     int       `gorm:"primaryKey"`
	Player       Player    `gorm:"foreignKey:PlayerID;association:ID;"`
	Time         time.Time `gorm:"type:TIME" json:"time,omitempty"`
	SeatID       int       `gorm:"default:NULL;"`
	Seat         Seat      `gorm:"foreignKey:SeatID;association:ID;"`
	HeroCard1ID  int       `gorm:"default:NULL;"`
	HeroCard1    Card      `gorm:"foreignKey:HeroCard1ID;association:ID;"`
	HeroCard2ID  int       `gorm:"default:NULL;"`
	HeroCard2    Card      `gorm:"foreignKey:HeroCard2ID;association:ID;"`
	TableCard1ID int       `gorm:"default:NULL;"`
	TableCard1   Card      `gorm:"foreignKey:TableCard1ID;association:ID;"`
	TableCard2ID int       `gorm:"default:NULL;"`
	TableCard2   Card      `gorm:"foreignKey:TableCard2ID;association:ID;"`
	TableCard3ID int       `gorm:"default:NULL;"`
	TableCard3   Card      `gorm:"foreignKey:TableCard3ID;association:ID;"`
	TableCard4ID int       `gorm:"default:NULL;"`
	TableCard4   Card      `gorm:"foreignKey:TableCard4ID;association:ID;"`
	TableCard5ID int       `gorm:"default:NULL;"`
	TableCard5   Card      `gorm:"foreignKey:TableCard5ID;association:ID;"`
	Gain         float64   `gorm:"type:float" json:"gain,omitempty"`
	PreFlopID    int       `gorm:"default:NULL;"`
	Preflop      Action    `gorm:"foreignKey:PreFlopID;association:ID;"`
	FlopID       int       `gorm:"default:NULL;"`
	Flop         Action    `gorm:"foreignKey:FlopID;association:ID;"`
	TurnID       int       `gorm:"default:NULL;"`
	Turn         Action    `gorm:"foreignKey:TurnID;association:ID;"`
	RiverID      int       `gorm:"default:NULL;"`
	River        Action    `gorm:"foreignKey:RiverID;association:ID;"`
}
