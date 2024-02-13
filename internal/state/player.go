package state

type Team string

const (
	NoTeam Team = "-"
	XTeam  Team = "X"
	OTeam  Team = "O"
)

type Player interface {
	GetTeam() Team
}

type HumanPlayer struct {
	team Team
}

func NewHumanPlayer(t Team) *HumanPlayer {
	return &HumanPlayer{
		team: t,
	}
}

func (h HumanPlayer) GetTeam() Team {
	return h.team
}
