package main

import (
	"fmt"

	utopiago "github.com/Sagleft/utopialib-go"
)

func (app *solution) onContactMessage(m utopiago.InstantMessage) {
	fmt.Println("[spy]")
}

func (app *solution) onChannelMessage(m utopiago.ChannelMessage) {
	if m.Text == "" {
		return
	}

	// TODO
}

func (app *solution) onPrivateChannelMessage(m utopiago.ChannelMessage) {
	fmt.Println("[PRIVATE] " + m.Nick + ": " + m.Text) // placeholder
}
