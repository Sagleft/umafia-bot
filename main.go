package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Sagleft/uchatbot-engine"
	coingecko "github.com/superoo7/go-gecko/v3"
)

const (
	configJSONPath             = "config.json"
	dataProviderConnectTimeout = 5 * time.Second
)

func main() {
	app := newSolution()

	err := checkErrors(
		app.parseConfig,
		app.dataProviderConnect,
		app.initBot,
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

func (app *solution) dataProviderConnect() error {
	app.DataProvider = coingecko.NewClient(&http.Client{
		Timeout: dataProviderConnectTimeout,
	})
	return nil
}

func (app *solution) getCryptonPrice() (float32, error) {
	data, err := app.DataProvider.SimpleSinglePrice("utopia", "usd")
	if err != nil {
		return 0, err
	}

	return data.MarketPrice, nil
}
