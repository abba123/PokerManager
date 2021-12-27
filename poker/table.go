package poker

import "time"

type Table struct {
	ID     int
	Time   time.Time
	Player map[string]Player
	Card   []Card
}
