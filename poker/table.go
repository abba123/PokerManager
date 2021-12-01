package poker

import "time"

type Table struct {
	ID     int
	Time   time.Time
	Player []Player
	Card   []Card
}
