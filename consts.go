package main

import "time"

const (
	configJSONPath                = "config.json"
	gameSessionNameLength         = 12
	dataProviderConnectTimeout    = 5 * time.Second
	sendChatMessagesBufferSize    = 200
	limitBotChatOneMessageTimeout = time.Millisecond * 1100
	dayDiscussionTime             = time.Second * 15

	botStartedMessage  = "Бот запущен\n[spy]"
	botStoppedMesssage = "Бот остановлен\n[spy]"
)

var (
	startGameCommands = map[string]struct{}{
		"играть":      {},
		"начать игру": {},
		"play":        {},
		"play game":   {},
		"start game":  {},
	}
)
