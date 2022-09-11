package game

import (
	"time"
)

// players vote to start game.
// returns true when min players found
func (s *Session) awaitStartVote() bool {
	// wait
	time.Sleep(startGameVoteDuration)

	// count votes
	if s.PlayersCount < minPlayersCount {
		s.narrator("Не набрано минимальное число игроков.\nОтмена старта игры")
		return false
	}

	return true
}
