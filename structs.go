package main

import (
	"github.com/Sagleft/uchatbot-engine"
	utopiago "github.com/Sagleft/utopialib-go"
	coingecko "github.com/superoo7/go-gecko/v3"
)

type solution struct {
	Engine       *uchatbot.ChatBot
	DataProvider *coingecko.Client
	Config       config
}

type config struct {
	Utopia utopiago.UtopiaClient `json:"utopia"`
	Chats  []uchatbot.Chat       `json:"chats"`
}
