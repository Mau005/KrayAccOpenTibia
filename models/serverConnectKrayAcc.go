package models

type ServerConnectKrayAcc struct {
	IpWebApi    string `yaml:"IpWebApi" json:"ip_web_api"`
	PortApi     uint   `yaml:"PortApi" json:"port_api"`
	Token       string
	ClientWorld []ClientWorld
}
