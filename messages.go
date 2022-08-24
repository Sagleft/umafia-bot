package main

import (
	"fmt"
	"strings"

	utopiago "github.com/Sagleft/utopialib-go"
)

func (app *solution) onContactMessage(m utopiago.InstantMessage) {
	fmt.Println("[CONTACT] " + m.Nick + ": " + m.Text) // placeholder
}

func (app *solution) onChannelMessage(m utopiago.ChannelMessage) {
	if m.Text == "" {
		return
	}

	if strings.ToLower(m.Text) == "crypton rate" {
		rate, err := app.getCryptonPrice()
		if err != nil {
			app.onError(err)
			return
		}

		msg := "1 [diamond] = " + formatFloat(float64(rate), 2) + "$"
		_, err = app.Config.Utopia.SendChannelMessage(m.ChannelID, msg)
		if err != nil {
			app.onError(err)
			return
		}
	}
}

func (app *solution) onPrivateChannelMessage(m utopiago.ChannelMessage) {
	fmt.Println("[PRIVATE] " + m.Nick + ": " + m.Text) // placeholder
}
