package main

import (
	"bot/game"
	"fmt"
	"log"
	"reflect"
	"strings"

	swissknife "github.com/Sagleft/swiss-knife"
	utopiago "github.com/Sagleft/utopialib-go"
)

// remove bot old messages
/*func (b *bot) removeBotMessages() error {
	for _, chat := range b.Config.Chats {
		messages, err := b.Config.Utopia.GetChannelMessages(chat.ID, 0, maxMessagesInOneDelete)
		if err != nil {
			return err
		}
		for i := 0; i < len(messages); i++ {
			msg := messages[i]
			if msg.Text == botStartedMessage || msg.Text == botStoppedMesssage {
				b.removeChatMessage(chat.ID, msg.ID)
			}
		}
	}
	return nil
}*/

// when a user in the personal contacts list sent a message
func (b *bot) onContactMessage(m utopiago.InstantMessage) {
	fmt.Println("[spy]")
}

func (b *bot) isGameSessionAlreadyStarted(channelID string) bool {
	_, isFound := b.Sessions[channelID]
	return isFound
}

// when in one of the chats (channels) someone sent a message
func (b *bot) onChannelMessage(m utopiago.WsChannelMessage) {
	if m.Text == "" {
		return
	}

	isSessionStartedInChannel := b.isGameSessionAlreadyStarted(m.ChannelID)

	if isPlayGameCommand(m.Text) {
		if isSessionStartedInChannel {
			b.sendChatMessage(m.ChannelID, "Игра уже запущена")
			return
		}
		b.startNewGameSession(m.ChannelID)
		return
	}

	if isSessionStartedInChannel {
		b.routeGameSessionMessage(m.ChannelID, m)
	}
}

func (b *bot) routeGameSessionMessage(channelID string, m utopiago.WsChannelMessage) {
	b.Sessions[channelID].HandleMessage(game.HandleMessageTask{
		Text:             m.Text,
		PlayerPubkeyHash: m.PubkeyHash,
		PlayerNickname:   m.Nick,
	})
}

// when someone sends a message in a chat private room section
func (b *bot) onPrivateChannelMessage(m utopiago.WsChannelMessage) {
	// TODO
}

// add message to chat queue
func (b *bot) sendChatMessage(channelID, message string) {
	b.Workers.Chat.W.AddEvent(postChatMessageTask{
		Text:      message,
		ChannelID: channelID,
	})
}

// add message to private chat queue
func (b *bot) sendPrivateMessage(channelID, pubkeyHash, message string) {
	b.Workers.Chat.W.AddEvent(privateChatMessageTask{
		Text:              message,
		ContactPubkeyHash: pubkeyHash,
		ChannelID:         channelID,
	})
}

func (b *bot) sendPrivateChatMessageFromQueue(event interface{}) {
	// get message
	message, isConvertable := event.(privateChatMessageTask)
	if !isConvertable {
		log.Println("invalid event received in channel worker: " + reflect.TypeOf(event).String())
		return
	}

	// sync messages rate
	b.Workers.PrivateChat.R.Wait()

	// send channel message
	_, err := b.Config.Utopia.SendChannelContactMessage(
		message.ChannelID,
		message.ContactPubkeyHash,
		message.Text,
	)
	if err != nil {
		log.Println(err)
	}
}

func (b *bot) sendChatMessageFromQueue(event interface{}) {
	// get message
	message, isConvertable := event.(postChatMessageTask)
	if !isConvertable {
		log.Println("invalid event received in channel worker: " + reflect.TypeOf(event).String())
		return
	}

	// sync messages rate
	b.Workers.Chat.R.Wait()

	// send channel message
	_, err := b.Config.Utopia.SendChannelMessage(message.ChannelID, message.Text)
	if err != nil {
		log.Println(err)
	}
}

// add message to delete queue
/*func (b *bot) removeChatMessage(channelID string, messageID int64) {
	b.Workers.RemoveMessages.W.AddEvent(removeChatMessageTask{
		ChannelID: channelID,
		MessageID: messageID,
	})
}*/

func (b *bot) removeChatMessageFromQueue(event interface{}) {
	// get message
	message, isConvertable := event.(removeChatMessageTask)
	if !isConvertable {
		log.Println("invalid event received in remove-messages worker: " + reflect.TypeOf(event).String())
		return
	}

	// sync messages rate
	b.Workers.RemoveMessages.R.Wait()

	// remove message
	err := b.Config.Utopia.RemoveChannelMessage(message.ChannelID, message.MessageID)
	if err != nil {
		log.Println(err)
	}
}

func (b *bot) startNewGameSession(channelID string) {
	b.Sessions[channelID] = game.NewSession(game.SessionData{
		Name:      strings.ToUpper(swissknife.GetRandomString(gameSessionNameLength)),
		ChannelID: channelID,
		Callbacks: game.SessionCallbacks{
			SendNarratorMessage:      b.sendNarratorMessage,
			SendPlayerPrivateMessage: b.sendPlayerPrivateMessage,
			RemoveSession:            b.removeSession,
		},
	})
	b.Sessions[channelID].Start()
}

func (b *bot) sendNarratorMessage(m game.SendNarratorMessageTask) {
	b.sendChatMessage(m.ChannelID, m.Message)
}

func (b *bot) sendPlayerPrivateMessage(m game.SendPlayerMessageTask) {
	b.sendPrivateMessage(m.ChannelID, m.PlayerPubkeyHash, m.Message)
}

func (b *bot) removeSession(channelID string) {
	if !b.isGameSessionAlreadyStarted(channelID) {
		return
	}

	log.Println("remove session from channel " + channelID)
	delete(b.Sessions, channelID)
}
