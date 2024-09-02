package models

type SolicitudeLayout struct {
	NavBar          bool
	Menu            bool
	Footer          bool
	TopPlayers      bool
	News            bool
	ServerStatus    bool
	Rates           bool
	Login           bool
	Modal           bool
	Scripts         bool
	WhoIsOnline     bool
	LastDeath       bool
	HighScore       bool
	Guilds          bool
	Staff           bool
	RecoveryAccount bool
	Dowloads        bool
	PoliticService  bool
}

type Layout struct {
	NavBar          string
	Menu            string
	Footer          string
	TopPlayers      string
	News            string
	ServerStatus    string
	Rates           string
	Login           string
	Modal           string
	Scripts         string
	WhoIsOnline     string
	LastDeath       string
	HighScore       string
	Guilds          string
	Staff           string
	RecoveryAccount string
	Dowloads        string
	PoliticService  string
	LogoButtons     string
	Head            string
}

func NewLayoutDefault() SolicitudeLayout {
	return SolicitudeLayout{Login: true, ServerStatus: true, TopPlayers: true, Rates: true}
}
