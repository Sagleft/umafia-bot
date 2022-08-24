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
	app := newSolution()

	err := checkErrors(
		app.parseConfig,
		app.initBot,
		app.initChannelWorkers,
	)
	if err != nil {
		log.Fatalln(err)
	}

	app.printLaunched()
	app.runInBackground()
}

func newSolution() *solution {
	return &solution{}
}

func (app *solution) printLaunched() {
	fmt.Println("bot started")
}

func (app *solution) runInBackground() {
	forever := make(chan bool)
	// run in background
	<-forever
}

func (app *solution) initBot() error {
	var err error
	app.Engine, err = uchatbot.NewChatBot(uchatbot.ChatBotData{
		Client: &app.Config.Utopia,
		Chats:  app.Config.Chats,
		Callbacks: uchatbot.ChatBotCallbacks{
			OnContactMessage:        app.onContactMessage,
			OnChannelMessage:        app.onChannelMessage,
			OnPrivateChannelMessage: app.onPrivateChannelMessage,
		},
		UseErrorCallback: true,
		ErrorCallback:    app.onError,
	})
	return err
}

func (app *solution) onError(err error) {
	log.Println(err)
}

type chatMessage struct {
	Text      string
	ChannelID string
}

func (app *solution) initChannelWorkers() error {
	app.Workers.ChatWorker = swissknife.NewChannelWorker(app.sendChatMessage, sendChatMessagesBufferSize)
	app.Workers.ChatMessagesLimiter = rate.New(1, limitBotChatOneMessageTimeout)
	return nil
}

func (app *solution) sendChatMessage(event interface{}) {
	// get message
	message, isConvertable := event.(chatMessage)
	if !isConvertable {
		log.Println("invalid event received in channel worker: " + reflect.TypeOf(event).String())
		return
	}

	// sync messages rate
	app.Workers.ChatMessagesLimiter.Wait()

	// send channel message
	_, err := app.Config.Utopia.SendChannelMessage(message.ChannelID, message.Text)
	if err != nil {
		log.Println(err)
	}
}
