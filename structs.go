package main

import (
	"os"

	swissknife "github.com/Sagleft/swiss-knife"
	"github.com/Sagleft/uchatbot-engine"
	utopiago "github.com/Sagleft/utopialib-go"
	"github.com/beefsack/go-rate"
)

type bot struct {
	Engine   *uchatbot.ChatBot
	Config   config
	Workers  channelWorkers
	Sessions gameSessions

	OnExit chan os.Signal
}

// channel ID -> session
type gameSessions map[string]*gameSession

type gameSession struct {
	ChannelID string
	//LangCode  string
}

type channelWorkers struct {
	ChatMessagesLimiter *rate.RateLimiter
	ChatWorker          *swissknife.ChannelWorker
}

type config struct {
	Utopia utopiago.UtopiaClient `json:"utopia"`
	Chats  []uchatbot.Chat       `json:"chats"`
}
