package game

/*
==================================================
                       BASE
==================================================
*/

type actorBase struct {
	Type         string
	PlayerInGame bool
	PlayerSleep  bool
}

func (a *actorBase) GetType() string {
	return a.Type
}

func (a *actorBase) InGame() bool {
	return a.PlayerInGame
}

func (a *actorBase) IsSleep() bool {
	return a.PlayerSleep
}

type actor interface {
	GetType() string
	InGame() bool
	IsSleep() bool

	Kill()
	Cure()
	Check()
}

/*
==================================================
                     ACTORS
==================================================
*/

type Mafia struct {
	actorBase
}

type Doctor struct {
	actorBase
}

type Civilian struct {
	actorBase
}

type Commissar struct {
	actorBase
}
