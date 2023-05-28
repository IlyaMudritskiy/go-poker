package objects

type PlayerRoles struct {
	Dealer     int
	SmallBlind int
	BigBlind   int
	FirstBet   int
}

var Roles = PlayerRoles{0, 1, 2, 3}
