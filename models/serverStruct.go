package models

type ServerStatus struct {
	IsOnline         bool
	PlayersCount     int
	PlayersMaxCount  int
	PlayersPeakCount int
	MapName          string
	MapAuthor        string
	MapWidth         int
	MapHeight        int
	NPCs             int
	Monsters         int
	Uptime           string
	Location         string
	URL              string
	Client           string
	Server           string
	ServerName       string
	ServerIP         string
	MOTD             string
}
