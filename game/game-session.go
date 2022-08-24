package game

import (
	"fmt"

	"github.com/SolarLune/gofsm"
)

func NewSession(data SessionData) *Session {
	s := &Session{
		FSM:  gofsm.NewFSM(),
		Data: data,
	}

	s.initStates()
	return s
}

func (s *Session) Start() {
	s.changeState(defaultState)
	fmt.Println("game session started in channel " + s.Data.ChannelID)
}
