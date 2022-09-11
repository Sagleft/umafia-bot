package game

import (
	"fmt"

	"github.com/SolarLune/gofsm"
)

func NewSession(data SessionData) *Session {
	s := &Session{
		FSM:     gofsm.NewFSM(),
		Data:    data,
		Players: make(playersMap),
	}

	s.initStates()
	return s
}

func (s *Session) Start() {
	fmt.Println("game session `" + s.Data.Name + "` started in channel " + s.Data.ChannelID)
	s.changeState(defaultState)
}

type HandleMessageTask struct {
	Text             string
	PlayerPubkeyHash string
	PlayerNickname   string
}

func (s *Session) HandleMessage(m HandleMessageTask) {
	switch s.FSM.State {
	case stateInit:
		s.routeInitMessage(m)
		return
	}
}

func (s *Session) routeInitMessage(m HandleMessageTask) {

}
