package game

import (
	"fmt"

	"github.com/SolarLune/gofsm"
)

func (s *Session) initStates() {
	s.FSM.Register(stateInit, gofsm.State{Update: s.onInit})
	s.FSM.Register(stateDay, gofsm.State{Update: s.onDay})
	s.FSM.Register(stateNight, gofsm.State{Update: s.onNight})
	s.FSM.Register(stateVote, gofsm.State{Update: s.onVote})
	s.FSM.Register(stateFinish, gofsm.State{Update: s.onFinish})
}

func (s *Session) changeState(newState string) {
	fmt.Println("change session `" + s.Data.Name + "` state to `" + newState + "`")
	s.Data.LastState = newState
	s.FSM.Change(newState)
}

func (s *Session) onInit() {}

func (s *Session) onDay() {}

func (s *Session) onVote() {}

func (s *Session) onNight() {}

func (s *Session) onFinish() {}

func (s *Session) goToInit() {
	s.changeState(stateInit)
}

func (s *Session) goToDay() {
	s.changeState(stateDay)
}

func (s *Session) goToVote() {
	s.changeState(stateVote)
}

func (s *Session) goToNight() {
	s.changeState(stateNight)
}

func (s *Session) goToFinish() {
	s.changeState(stateFinish)
}
