package main

import (
	"bot/game"
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
type gameSessions map[string]*game.Session

type botWorker struct {
	R *rate.RateLimiter
	W *swissknife.ChannelWorker
}

type channelWorkers struct {
	Chat           *botWorker
	PrivateChat    *botWorker
	RemoveMessages *botWorker
}

type config struct {
	Utopia utopiago.UtopiaClient `json:"utopia"`
	Chats  []uchatbot.Chat       `json:"chats"`
}

type privateChatMessageTask struct {
	Text              string
	ContactPubkeyHash string
	ChannelID         string
}

type postChatMessageTask struct {
	Text      string
	ChannelID string
}

type removeChatMessageTask struct {
	ChannelID string
	MessageID int64
}
