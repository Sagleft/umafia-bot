package main

import (
	"github.com/Sagleft/uchatbot-engine"
	utopiago "github.com/Sagleft/utopialib-go"
)

type solution struct {
	Engine *uchatbot.ChatBot
	Config config
}

type config struct {
	Utopia utopiago.UtopiaClient `json:"utopia"`
	Chats  []uchatbot.Chat       `json:"chats"`
}
