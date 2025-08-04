package models

type NavWeb struct {
	Authentication bool
	AccountID      int
	TypeAccess     int
	IsPremmium     bool
	AccountName    string
	MyPlayers      []Players
	//Guilds!
}
