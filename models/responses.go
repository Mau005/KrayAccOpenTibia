package models

type PlayData struct {
	World      []ClientWorld      `json:"worlds"`
	Characters []ClientCharacters `json:"characters"`
}

type ResponseData struct {
	Session  ClientSession `json:"session"`
	PlayData PlayData      `json:"playdata"`
}
