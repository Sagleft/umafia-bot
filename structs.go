package main

import (
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
}

// channel ID -> session
type gameSessions map[string]*gameSession

type gameSession struct{}

type channelWorkers struct {
	ChatMessagesLimiter *rate.RateLimiter
	ChatWorker          *swissknife.ChannelWorker
}

type config struct {
	Utopia utopiago.UtopiaClient `json:"utopia"`
	Chats  []uchatbot.Chat       `json:"chats"`
}
