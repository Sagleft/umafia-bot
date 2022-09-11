package game

import "github.com/SolarLune/gofsm"

type Session struct {
	FSM     *gofsm.FSM
	Data    SessionData
	Players playersMap
}

type playerHash string

type playerData struct {
	Nick  string
	Hash  playerHash // pubkey hash
	Actor actor
}

type playersMap map[playerHash]*playerData

type SendNarratorMessageTask struct {
	ChannelID string
	Message   string
}

type SendPlayerMessageTask struct {
	ChannelID        string
	PlayerPubkeyHash string
	Message          string
}

type SessionCallbacks struct {
	SendNarratorMessage      func(SendNarratorMessageTask)
	SendPlayerPrivateMessage func(SendPlayerMessageTask)
	RemoveSession            func(channelID string)
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
