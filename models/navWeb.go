package models

type NavWeb struct {
	Authentication bool
	AccountID      int
	TypeAccess     int
	IsPremmium     bool
	MyPlayers      []Players
	//Guilds!
}
