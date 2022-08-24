package game

import "github.com/SolarLune/gofsm"

type Session struct {
	FSM *gofsm.FSM

	Data SessionData
}

type SessionData struct {
	ChannelID string `json:"channelID"`
	DayNumber int    `json:"dayNumber"`
}
