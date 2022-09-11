package game

import "time"

const (
	// fsm states
	stateInit    = "init"
	stateDay     = "day"
	stateNight   = "night"
	stateVote    = "vote"
	stateFinish  = "finish"
	stateClear   = "clear"
	defaultState = stateInit

	startGameVoteDuration = time.Second * 30
	minPlayersCount       = 4
)
