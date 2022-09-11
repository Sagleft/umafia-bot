package game

func (s *Session) narrator(message string) {
	s.Data.Callbacks.SendNarratorMessage(SendNarratorMessageTask{
		ChannelID: s.Data.ChannelID,
		Message:   message,
	})
}

// inform player in channel private room
func (s *Session) informPlayer(hash playerHash, message string) {
	s.Data.Callbacks.SendPlayerPrivateMessage(SendPlayerMessageTask{
		ChannelID:        s.Data.ChannelID,
		PlayerPubkeyHash: string(hash),
		Message:          message,
	})
}

type HandleMessageTask struct {
	Text             string
	PlayerPubkeyHash playerHash
	PlayerNickname   string
}

func (s *Session) HandleMessage(m HandleMessageTask) {
	switch s.FSM.State {
	case stateInit:
		s.routeInitMessage(m)
		return
	case stateNight:
		s.routeNightMessage(m)
		return
	}
}

// INIT
func (s *Session) routeInitMessage(m HandleMessageTask) {
	if m.Text == "+" {
		if s.isPlayerJoined(m.PlayerPubkeyHash) {
			return // ignore join message duplicate
		}

		// add player
		s.addPlayer(playerData{
			Nick: m.PlayerNickname,
			Hash: playerHash(m.PlayerPubkeyHash),
		})
	}
}

// NIGHT
func (s *Session) routeNightMessage(m HandleMessageTask) {
	player := s.getPlayer(m.PlayerPubkeyHash)
	if player == nil {
		// user not found in game session
		return
	}

	if !player.Actor.InGame() {
		s.informPlayer(player.Hash, "Ты выбыл(а).\nПодожди до следующей игры")
		return
	}
}
