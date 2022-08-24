package main

import (
	"fmt"
	"log"
	"reflect"

	utopiago "github.com/Sagleft/utopialib-go"
)

// when a user in the personal contacts list sent a message
func (b *bot) onContactMessage(m utopiago.InstantMessage) {
	fmt.Println("[spy]")
}

func (b *bot) isGameSessionAlreadyStarted(channelID string) bool {
	_, isFound := b.Sessions[channelID]
	return isFound
}

// when in one of the chats (channels) someone sent a message
func (b *bot) onChannelMessage(m utopiago.ChannelMessage) {
	if m.Text == "" {
		return
	}

	if isPlayGameCommand(m.Text) {
		if b.isGameSessionAlreadyStarted(m.ChannelID) {
			b.sendChatMessage(m.ChannelID, "Игра уже запущена")
			return
		}
		b.startNewGameSession(m.ChannelID)
		return
	}

	// TODO
}

func (b *bot) startNewGameSession(channelID string) {
	b.Sessions[channelID] = &gameSession{
		ChannelID: channelID,
	}
}

// when someone sends a message in a chat private room section
func (b *bot) onPrivateChannelMessage(m utopiago.ChannelMessage) {
	// TODO
}

// add message to chat queue
func (b *bot) sendChatMessage(channelID, message string) {
	b.Workers.ChatWorker.AddEvent(chatMessage{
		Text:      message,
		ChannelID: channelID,
	})
}

func (b *bot) sendChatMessageFromQueue(event interface{}) {
	// get message
	message, isConvertable := event.(chatMessage)
	if !isConvertable {
		log.Println("invalid event received in channel worker: " + reflect.TypeOf(event).String())
		return
	}

	// sync messages rate
	b.Workers.ChatMessagesLimiter.Wait()

	// send channel message
	_, err := b.Config.Utopia.SendChannelMessage(message.ChannelID, message.Text)
	if err != nil {
		log.Println(err)
	}
}
