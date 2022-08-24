package main

import "time"

const (
	configJSONPath                = "config.json"
	dataProviderConnectTimeout    = 5 * time.Second
	sendChatMessagesBufferSize    = 200
	limitBotChatOneMessageTimeout = time.Millisecond * 1100
)
