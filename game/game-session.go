package game

import (
	"fmt"

	"github.com/SolarLune/gofsm"
)

func NewSession(data SessionData) *Session {
	return &Session{
		FSM:  gofsm.NewFSM(),
		Data: data,
	}
}

func (g *Session) Start() {
	// TODO

	fmt.Println("game session started in channel " + g.Data.ChannelID)
}
