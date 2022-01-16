package poker

type Player struct {
	Name   string  
	Seat   string  
	Gain   float64 
	Action struct {
		Preflop string
		Flop    string
		Turn    string
		River   string
	} 
	Card      []Card 
	Rank      int    
	RankValue []Card 
}
