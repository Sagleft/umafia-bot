package game

func (s *Session) narrator(message string) {
	s.Data.Callbacks.SendNarratorMessage(SendNarratorMessageTask{
		ChannelID: s.Data.ChannelID,
		Message:   message,
	})
}

// inform player in channel private room
func (s *Session) informPlayer(playerHash string, message string) {
	s.Data.Callbacks.SendPlayerPrivateMessage(SendPlayerMessageTask{
		ChannelID:        s.Data.ChannelID,
		PlayerPubkeyHash: playerHash,
		Message:          message,
	})
}
