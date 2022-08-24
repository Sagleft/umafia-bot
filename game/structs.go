package game

import "github.com/SolarLune/gofsm"

type Session struct {
	FSM  *gofsm.FSM
	Data SessionData
}

type SessionCallbacks struct {
	SendGameStartMessage func(*SessionData)
}

type SessionData struct {
	// required
	Name      string           `json:"name"`
	ChannelID string           `json:"channelID"`
	Callbacks SessionCallbacks `json:"-"`

	// optional
	DayNumber int    `json:"dayNumber"`
	LastState string `json:"lastState"`
}
