package game

/*
==================================================
                       BASE
==================================================
*/

type actorBase struct {
	RoleType     string
	RoleName     string
	AboutMessage string
	PlayerInGame bool
	PlayerSleep  bool
}

func (a *actorBase) GetType() string {
	return a.RoleType
}

func (a *actorBase) GetRoleName() string {
	return a.RoleName
}

func (a *actorBase) GetAboutMessage() string {
	return a.AboutMessage
}

func (a *actorBase) InGame() bool {
	return a.PlayerInGame
}

func (a *actorBase) IsSleep() bool {
	return a.PlayerSleep
}

type actor interface {
	GetRoleName() string
	GetAboutMessage() string
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
			RoleType: "mafia",
			RoleName: "мафиози",
			AboutMessage: "Во время ночного голосования пришли мне команду для выбора цели мафии.\n\n" +
				"Пример:\nубить 11\n\nгде 11 - игрок под номером 11",
			PlayerInGame: true,
		},
	}
}

func newDoctorPlayer() *Doctor {
	return &Doctor{
		actorBase: actorBase{
			RoleType: "doctor",
			RoleName: "доктор",
			AboutMessage: "Ночью пришли мне команду для выбора кого вылечить.\n\n" +
				"Пример:\nлечить 11\n\nгде 11 - игрок, которого надо вылечить",
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
			RoleType: "commissar",
			RoleName: "комиссар",
			AboutMessage: "Ночью выбери одну из команд и напиши мне.\n\n" +
				"Примеры:\nубить 11 - застрелить игрока под номером 11.\n" +
				"узнать 12 - узнать роль игрока под номером 12",
			PlayerInGame: true,
		},
	}
}
