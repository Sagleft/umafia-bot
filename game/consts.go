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
	dayTalkDuration       = time.Second * 5
	voteDuration          = time.Second * 20
	nightDuration         = time.Second * 30
)
