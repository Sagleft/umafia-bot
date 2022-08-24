package main

import (
	"fmt"
	"log"
	"reflect"

	utopiago "github.com/Sagleft/utopialib-go"
)

func (b *bot) onContactMessage(m utopiago.InstantMessage) {
	fmt.Println("[spy]")
}

func (b *bot) onChannelMessage(m utopiago.ChannelMessage) {
	if m.Text == "" {
		return
	}

	// TODO
}

func (b *bot) onPrivateChannelMessage(m utopiago.ChannelMessage) {
	// TODO
}

func (b *bot) sendChatMessage(event interface{}) {
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
