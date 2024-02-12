package state

type StateChange string

const (
	TeamSelected StateChange = "T"
	BoardChanged StateChange = "B"
	GameOvered   StateChange = "O"
)

type Statement struct {
	State  StateChange
	Team   Team
	Square Square
}

func NewStatement(e StateChange, t Team, s Square) *Statement {
	return &Statement{
		State:  e,
		Team:   t,
		Square: s,
	}
}

func NewBoardStatement(t Team, s Square) *Statement {
	return NewStatement(BoardChanged, t, s)
}

//TODO validate
