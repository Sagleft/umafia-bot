package game

import (
	"fmt"
	"log"

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

func (s *Session) isPlayerJoined(playerPubkeyHash string) bool {
	_, isJoined := s.Players[playerHash(playerPubkeyHash)]
	return isJoined
}

func (s *Session) getPlayersCount() int {
	return len(s.Players)
}

func (s *Session) addPlayer(d playerData) {
	log.Println("add player " + d.Nick + " to game in " + s.Data.ChannelID)
	s.Players[d.Hash] = &d
}
