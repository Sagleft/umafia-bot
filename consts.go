package main

import "time"

const (
	configJSONPath                    = "config.json"
	gameSessionNameLength             = 12
	dataProviderConnectTimeout        = 5 * time.Second
	sendChatMessagesBufferSize        = 100
	sendPrivateChatMessagesBufferSize = 100
	deleteMessagesBufferSize          = 100
	limitBotChatOneMessageTimeout     = time.Millisecond * 1100
	limitBotContactMessageTimeout     = time.Millisecond * 1100
	limitChatMessageDeleteTimeout     = time.Millisecond * 1100
	maxMessagesInOneDelete            = 15
	dayDiscussionTime                 = time.Second * 15

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
		"запуск":      {},
		"старт":       {},
	}
)
