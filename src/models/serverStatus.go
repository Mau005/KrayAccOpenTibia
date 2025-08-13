package models

import "encoding/xml"

type ServerStatus struct {
	XMLName    xml.Name `xml:"tsqp"`
	Version    string   `xml:"version,attr"`
	ServerInfo struct {
		Uptime     string `xml:"uptime,attr"`
		IP         string `xml:"ip,attr"`
		ServerName string `xml:"servername,attr"`
		Port       string `xml:"port,attr"`
		Location   string `xml:"location,attr"`
		URL        string `xml:"url,attr"`
		Server     string `xml:"server,attr"`
		Version    string `xml:"version,attr"`
		Client     string `xml:"client,attr"`
	} `xml:"serverinfo"`
	Owner struct {
		Name  string `xml:"name,attr"`
		Email string `xml:"email,attr"`
	} `xml:"owner"`
	Players struct {
		Online int `xml:"online,attr"`
		Max    int `xml:"max,attr"`
		Peak   int `xml:"peak,attr"`
	} `xml:"players"`
	Monsters struct {
		Total int `xml:"total,attr"`
	} `xml:"monsters"`
	NPCs struct {
		Total int `xml:"total,attr"`
	} `xml:"npcs"`
	Rates struct {
		Experience int `xml:"experience,attr"`
		Skill      int `xml:"skill,attr"`
		Loot       int `xml:"loot,attr"`
		Magic      int `xml:"magic,attr"`
		Spawn      int `xml:"spawn,attr"`
	} `xml:"rates"`
	Map struct {
		Name   string `xml:"name,attr"`
		Author string `xml:"author,attr"`
		Width  int    `xml:"width,attr"`
		Height int    `xml:"height,attr"`
	} `xml:"map"`
	MOTD string `xml:"motd"`
}
