package game

/*
==================================================
                       BASE
==================================================
*/

type actorBase struct {
	RoleType     string
	RoleName     string
	PlayerInGame bool
	PlayerSleep  bool
}

func (a *actorBase) GetType() string {
	return a.RoleType
}

func (a *actorBase) GetRoleName() string {
	return a.RoleName
}

func (a *actorBase) InGame() bool {
	return a.PlayerInGame
}

func (a *actorBase) IsSleep() bool {
	return a.PlayerSleep
}

type actor interface {
	GetRoleName() string
	GetType() string
	InGame() bool
	IsSleep() bool

	//Kill()
	//Cure()
	//Check()
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

func newMafiaPlayer() *Mafia {
	return &Mafia{
		actorBase: actorBase{
			RoleType:     "mafia",
			RoleName:     "мафиози",
			PlayerInGame: true,
		},
	}
}

func newDoctorPlayer() *Doctor {
	return &Doctor{
		actorBase: actorBase{
			RoleType:     "doctor",
			RoleName:     "доктор",
			PlayerInGame: true,
		},
	}
}

func newCivilianPlayer() *Civilian {
	return &Civilian{
		actorBase: actorBase{
			RoleType:     "civilian",
			RoleName:     "мирный",
			PlayerInGame: true,
		},
	}
}

func newCommissarPlayer() *Commissar {
	return &Commissar{
		actorBase: actorBase{
			RoleType:     "commissar",
			RoleName:     "комиссар",
			PlayerInGame: true,
		},
	}
}
