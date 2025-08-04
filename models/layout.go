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
	MyAccount       bool
	Discord         bool
}

type Layout struct {
	//Standar web
	NavBar    string
	Component string

	//Configuration WEB
	Modal       string
	Scripts     string
	LogoButtons string
	Head        string

	//SideBar
	Login        string
	Footer       string
	TopPlayers   string
	ServerStatus string
	Rates        string
	Discord      string
}

func NewLayoutDefault() SolicitudeLayout {
	return SolicitudeLayout{Login: true, ServerStatus: true, TopPlayers: true, Rates: true, Discord: true}
}
