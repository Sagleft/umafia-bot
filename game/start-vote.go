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
	if s.getPlayersCount() < minPlayersCount {
		s.narrator("Не набрано минимальное число игроков.\nОтмена старта игры\n" + ":^|")
		return false
	}

	return true
}
