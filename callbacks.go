package main

import "bot/game"

func (b *bot) onGameSessionStarted(d *game.SessionData) {
	b.sendChatMessage(d.ChannelID, "Начинаем голосование на запуск игры")
}
