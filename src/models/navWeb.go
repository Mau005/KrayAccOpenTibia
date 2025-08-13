package models

type NavWeb struct {
	Authentication bool
	AccountID      int
	TypeAccess     uint8
	IsPremmium     bool
	AccountName    string
	MyPlayers      []Players
	//Guilds!
}
