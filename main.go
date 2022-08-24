package main

import (
	"fmt"
	"log"
	"reflect"

	swissknife "github.com/Sagleft/swiss-knife"
	"github.com/Sagleft/uchatbot-engine"
	"github.com/beefsack/go-rate"
)

func main() {
	b := newSolution()

	err := checkErrors(
		b.parseConfig,
		b.initBot,
		b.initChannelWorkers,
	)
	if err != nil {
		log.Fatalln(err)
	}

	b.printLaunched()
	b.runInBackground()
}

func newSolution() *bot {
	return &bot{}
}

func (b *bot) printLaunched() {
	fmt.Println("bot started")
}

func (b *bot) runInBackground() {
	forever := make(chan bool)
	// run in background
	<-forever
}

func (b *bot) initBot() error {
	var err error
	b.Engine, err = uchatbot.NewChatBot(uchatbot.ChatBotData{
		Client: &b.Config.Utopia,
		Chats:  b.Config.Chats,
		Callbacks: uchatbot.ChatBotCallbacks{
			OnContactMessage:        b.onContactMessage,
			OnChannelMessage:        b.onChannelMessage,
			OnPrivateChannelMessage: b.onPrivateChannelMessage,
		},
		UseErrorCallback: true,
		ErrorCallback:    b.onError,
	})
	return err
}

func (b *bot) onError(err error) {
	log.Println(err)
}

type chatMessage struct {
	Text      string
	ChannelID string
}

func (b *bot) initChannelWorkers() error {
	b.Workers.ChatWorker = swissknife.NewChannelWorker(b.sendChatMessage, sendChatMessagesBufferSize)
	b.Workers.ChatMessagesLimiter = rate.New(1, limitBotChatOneMessageTimeout)
	return nil
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
