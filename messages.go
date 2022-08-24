package main

import (
	"fmt"

	utopiago "github.com/Sagleft/utopialib-go"
)

func (app *bot) onContactMessage(m utopiago.InstantMessage) {
	fmt.Println("[spy]")
}

func (app *bot) onChannelMessage(m utopiago.ChannelMessage) {
	if m.Text == "" {
		return
	}

	// TODO
}

func (app *bot) onPrivateChannelMessage(m utopiago.ChannelMessage) {
	// TODO
}
