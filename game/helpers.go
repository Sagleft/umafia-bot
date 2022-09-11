package game

func (s *Session) isItFirstDay() bool {
	return s.Data.DayNumber == 1
}

func (s *Session) getPlayer(hash playerHash) *playerData {
	player, isFound := s.Players[hash]
	if !isFound {
		return nil
	}

	return player
}
