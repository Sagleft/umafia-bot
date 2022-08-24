package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	swissknife "github.com/Sagleft/swiss-knife"
	"github.com/beefsack/go-rate"
)

type errorFunc func() error

func checkErrors(errChecks ...errorFunc) error {
	for _, errFunc := range errChecks {
		err := errFunc()
		if err != nil {
			return err
		}
	}
	return nil
}

func (app *bot) parseConfig() error {
	// parse config file
	if _, err := os.Stat(configJSONPath); os.IsNotExist(err) {
		return errors.New("failed to find config file")
	}

	jsonBytes, err := ioutil.ReadFile(configJSONPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonBytes, &app.Config)
}

func formatFloat(val float64, precision int) string {
	return strconv.FormatFloat(val, 'f', precision, 32)
}

func isPlayGameCommand(m string) bool {
	_, isFound := startGameCommands[m]
	return isFound
}

type getChannelWorkerTask struct {
	RateLimit  time.Duration
	Callback   func(event interface{})
	BufferSize int
}

func (b *bot) getChannelWorker(task getChannelWorkerTask) *botWorker {
	worker := &botWorker{
		R: rate.New(1, task.RateLimit),
		W: swissknife.NewChannelWorker(task.Callback, task.BufferSize),
	}

	go worker.W.Start()
	return worker
}
