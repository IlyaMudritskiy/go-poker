package game

// Pre-flop - Deal hands to players, bets
// Flop - Get 3 cards on table, bets
// River - Add 1 card to table, bets
// Turn - Add 1 card to table, bets, open cards

type Phase struct {
	Id int
}

type Game struct {
	Phase   [5]int
	Current int
}
