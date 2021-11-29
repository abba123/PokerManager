package poker

type Player struct {
	Name string
	Preflop int
	Flop int
	Turn int
	River int
	Card []Card
	Rank int
	RankValue []Card
}
