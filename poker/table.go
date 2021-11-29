package poker

import (
	"time"
)

type Table struct{
	Time time.Time
	HeroGain float32
	Player []Player
	Card []Card
}
