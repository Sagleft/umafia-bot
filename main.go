package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sagleft/uchatbot-engine"
)

func main() {
	b := newSolution()

	err := checkErrors(
		b.parseConfig,
		b.utopiaConnect,
		b.initChannelWorkers,
		b.notifyStarted,
	)
	if err != nil {
		log.Fatalln(err)
	}

	b.printLaunched()
	b.runInBackground()
}

func newSolution() *bot {
	b := &bot{
		Sessions: make(gameSessions),
		OnExit:   make(chan os.Signal, 1),
	}

	signal.Notify(b.OnExit, os.Interrupt, syscall.SIGTERM)
	go b.waitFinish()
	return b
}

func (b *bot) printLaunched() {
	fmt.Println("bot started")
}

func (b *bot) waitFinish() {
	<-b.OnExit

	fmt.Println()
	fmt.Println("stop bot..")
	b.notifyStopped()
	time.Sleep(time.Second * 2)

	os.Exit(1)
}

func (b *bot) runInBackground() {
	forever := make(chan bool)
	// run in background
	<-forever
}

func (b *bot) utopiaConnect() error {
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

func (b *bot) initChannelWorkers() error {
	b.Workers.ChatWorker = b.getChannelWorker(getChannelWorkerTask{
		RateLimit:  limitBotChatOneMessageTimeout,
		Callback:   b.sendChatMessageFromQueue,
		BufferSize: sendChatMessagesBufferSize,
	})
	return nil
}

func (b *bot) notifyStarted() error {
	b.notifyChats(botStartedMessage)
	return nil
}

func (b *bot) notifyStopped() {
	b.notifyChats(botStoppedMesssage)
}

func (b *bot) notifyChats(message string) {
	for _, chat := range b.Config.Chats {
		b.sendChatMessage(chat.ID, message)
	}
}
