package game

import "github.com/SolarLune/gofsm"

type Session struct {
	FSM  *gofsm.FSM
	Data SessionData

	PlayersCount int
	Players      playersMap
}

type playersMap map[string]struct{}

type SessionCallbacks struct {
	SendNarratorMessage func(*SessionData, string)
	RemoveSession       func(channelID string)
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
