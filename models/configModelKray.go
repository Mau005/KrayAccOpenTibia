package models

type ServerWeb struct {
	IP                   string `yaml:"IP"`
	Port                 uint16 `yaml:"Port"`
	Debug                bool   `yaml:"Debug"`
	ApiMode              bool   `yaml:"ApiMode"`
	LengthSecurity       int    `yaml:"LengthSecurity"`
	EnvironmentVariables bool   `yaml:"EnvironmentVariables"`
	UrlItemView          string `yaml:"UrlItemView"`
	UrlOutfitsView       string `yaml:"UrlOutfitsView"`
	TargetServer         string `yaml:"TargetServer"`
	LimitCreateCharacter uint8  `yaml:"LimitCreateCharacter"`
}

type MySQL struct {
	Host       string `yaml:"Host"`
	Port       uint16 `yaml:"Port"`
	UserName   string `yaml:"UserName"`
	DBPassword string `yaml:"DBPassword"`
	DataBase   string `yaml:"DataBase"`
}

type Certificate struct {
	ProtocolTLS bool   `yaml:"ProtocolHTTPS"`
	Chain       string `yaml:"Chain"`
	PrivKey     string `yaml:"PrivKey"`
}

type Configuration struct {
	DB          MySQL        `yaml:"MySQL"`
	ServerWeb   ServerWeb    `yaml:"ServerWeb"`
	Certificate Certificate  `yaml:"Certificate"`
	PoolSerer   []PoolServer `yaml:"ApiConnectionPool"`
}

//Local Server

type ExperienceStage struct {
	MinLevel   uint16
	MaxLevel   uint16
	Multiplier uint16
}

type RateServer struct {
	RateExp   uint16
	RateSkill uint16
	RateLoot  uint16
	RateMagic uint16
	RateSpawn uint16
}

type PoolServer struct {
	IpWebApi        string `yaml:"IpWebApi" json:"ip_web_api"`
	PortApi         uint   `yaml:"PortApi" json:"port_api"`
	Token           string
	World           ClientWorld
	RateServer      RateServer
	ExperienceStage ExperienceStage
}
