package game

import (
	"fmt"
	"strconv"

	"github.com/SolarLune/gofsm"
)

func (s *Session) initStates() {
	s.FSM.Register(stateInit, gofsm.State{Enter: s.onEnterInit, Update: s.onInit})
	s.FSM.Register(stateDay, gofsm.State{Enter: s.onEnterDay, Update: s.onDay})
	s.FSM.Register(stateNight, gofsm.State{Enter: s.onEnterNight, Update: s.onNight})
	s.FSM.Register(stateVote, gofsm.State{Enter: s.onEnterVote, Update: s.onVote})
	s.FSM.Register(stateFinish, gofsm.State{Enter: s.onEnterFinish, Update: s.onFinish})
	s.FSM.Register(stateClear, gofsm.State{Enter: s.onEnterClear, Update: s.onClear})
}

func (s *Session) changeState(newState string) {
	fmt.Println("change session `" + s.Data.Name + "` state to `" + newState + "`")
	s.Data.LastState = newState
	s.FSM.Change(newState)
}

func (s *Session) onEnterInit() {
	s.FSM.Update()
}

func (s *Session) narrator(message string) {
	s.Data.Callbacks.SendNarratorMessage(&s.Data, message)
}

func (s *Session) onInit() {
	s.narrator("Начинаем голосование на запуск игры.\n" +
		"Минимальное число игроков: " + strconv.Itoa(minPlayersCount) + ".\n" +
		"Жду " + strconv.Itoa(int(startGameVoteDuration.Seconds())) + " секунд")

	minPlayersFound := s.awaitStartVote()
	if !minPlayersFound {
		s.goToClear()
		return
	}

	s.onEnterDay()
}

func (s *Session) onEnterDay() {
	defer s.FSM.Update()
	s.narrator("ИГРА НАЧАЛАСЬ\n\nНаступил день")
}

func (s *Session) onDay() {}

func (s *Session) onEnterVote() {
	s.FSM.Update()
}

func (s *Session) onVote() {}

func (s *Session) onEnterNight() {
	s.FSM.Update()
}

func (s *Session) onNight() {}

func (s *Session) onEnterFinish() {
	s.FSM.Update()
}

func (s *Session) onFinish() {}

func (s *Session) onEnterClear() {
	s.FSM.Update()
}

func (s *Session) onClear() {
	s.Data.Callbacks.RemoveSession(s.Data.ChannelID)
}

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

func (s *Session) goToClear() {
	s.changeState(stateClear)
}
