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
