package models

import "time"

type PlayData struct {
	World      []ClientWorld      `json:"worlds"`
	Characters []ClientCharacters `json:"characters"`
}

type ResponseData struct {
	Session  ClientSession `json:"session"`
	PlayData PlayData      `json:"playdata"`
}

type Exception struct {
	Msg     string    `json:"msg"`
	TimeNow time.Time `json:"time_now"`
}
