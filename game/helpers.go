package game

func (g *Session) isItFirstDay() bool {
	return g.Data.DayNumber == 1
}
