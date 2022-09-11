package game

import (
	"fmt"
	"strconv"
	"time"

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

func (s *Session) onInit() {
	s.narrator("Начинаем голосование на запуск игры.\n" +
		"Минимальное число игроков: " + strconv.Itoa(minPlayersCount) + ".\n" +
		"Жду " + strconv.Itoa(int(startGameVoteDuration.Seconds())) + " секунд\n\n" +
		"Пиши + чтобы участвовать")

	minPlayersFound := s.awaitStartVote()
	if !minPlayersFound {
		s.goToClear()
		return
	}

	s.narrator("ИГРА НАЧАЛАСЬ\n\nРоли разосланы в приватные сообщения, прочти, чтобы узнать свою роль")
	s.goToDay()
}

func (s *Session) onEnterDay() {
	defer s.FSM.Update()
	s.narrator("Город просыпается")
}

func (s *Session) onDay() {
	if s.isItFirstDay() {
		s.setPlayerRoles()
		s.goToNight()
		return
	}

	// TODO: подведение итогов ночи

	time.Sleep(dayTalkDuration)

	s.goToVote()
}

func (s *Session) getRoles(playersCount int) []actor {
	roles := []actor{}
	for i := 0; i < playersCount; i++ {
		roles = append(roles, newCivilianPlayer())
	}

	// TODO: задать роли из соотношений на основе числа игроков

	return roles
}

func (s *Session) setPlayerRoles() {
	roles := s.getRoles(len(s.Players))

	// assign roles to players
	var i int = 0
	for _, player := range s.Players {

		// assign role
		player.Actor = roles[i]
		i++

		// notify player about his role
		// TODO: combine with mafia notify
		msg := "Твоя роль: " + player.Actor.GetRoleName() + "\n\n" + player.Actor.GetAboutMessage()
		s.informPlayer(player.Hash, msg)
	}
}

func (s *Session) onEnterVote() {
	defer s.FSM.Update()
	s.narrator("Начинаем голосование")
}

func (s *Session) onVote() {
	time.Sleep(voteDuration)

	// TODO: определение голосов

	s.goToNight()
}

func (s *Session) onEnterNight() {
	defer s.FSM.Update()

	s.narrator("Наступила ночь. Город засыпает")
}

func (s *Session) onNight() {
	// TODO: разные типы игроков делают свой выбор

	time.Sleep(nightDuration)

	s.goToDay()
}

func (s *Session) onEnterFinish() {
	defer s.FSM.Update()
	s.narrator("Игра завершена")
}

func (s *Session) onFinish() {
	s.goToClear()
}

func (s *Session) onEnterClear() {
	s.FSM.Update()
}

func (s *Session) onClear() {
	s.Data.Callbacks.RemoveSession(s.Data.ChannelID)
}

func (s *Session) goToNight() {
	s.changeState(stateNight)
}

func (s *Session) goToDay() {
	s.changeState(stateDay)
}

func (s *Session) goToVote() {
	s.changeState(stateVote)
}

func (s *Session) goToFinish() {
	s.changeState(stateFinish)
}

func (s *Session) goToClear() {
	s.changeState(stateClear)
}
