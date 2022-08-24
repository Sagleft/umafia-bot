package main

import (
	swissknife "github.com/Sagleft/swiss-knife"
	"github.com/Sagleft/uchatbot-engine"
	utopiago "github.com/Sagleft/utopialib-go"
	"github.com/beefsack/go-rate"
)

type solution struct {
	Engine  *uchatbot.ChatBot
	Config  config
	Workers channelWorkers
}

type channelWorkers struct {
	ChatMessagesLimiter *rate.RateLimiter
	ChatWorker          *swissknife.ChannelWorker
}

type config struct {
	Utopia utopiago.UtopiaClient `json:"utopia"`
	Chats  []uchatbot.Chat       `json:"chats"`
}
