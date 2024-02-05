package server

import (
	"errors"

	"github.com/google/uuid"
)

type GameState byte

const (
	GameCreated GameState = iota
	GameStarted
	GameOver
)

// Statement is the postion played as following:
// 0 | 1 | 2
// ---------
// 3 | 4 | 5
// ---------
// 6 | 7 | 8
type Statement byte

const (
	Statement0 Statement = iota
	Statement1
	Statement2
	Statement3
	Statement4 // middle
	Statement5
	Statement6
	Statement7
	Statement8
)

var ErrCanPlayTwice = errors.New("can not play twice in a row")
var ErrTeamAlreadySelected = errors.New("team already selected")
var ErrCantJoinStartedGame = errors.New("can not join started game")
var ErrWaitingPlayerJoin = errors.New("waiting player 2 join the game")
var ErrGameOverAlready = errors.New("game over already, somebody got fineshed!")
var ErrInvalidStatement = errors.New("statement invalid by brother")

type round struct {
	player    Player
	statement Statement
}

type Game struct {
	id    uuid.UUID
	state GameState

	// store round history
	lastRound  *round
	rounds     [9]*round
	roundsLeft byte

	// players of the game
	player1 Player
	player2 Player
}

func NewGame(p1 Player) *Game {
	return &Game{
		id:         uuid.New(),
		state:      GameCreated,
		player1:    p1,
		roundsLeft: 9,
	}
}

func (g *Game) Join(p2 Player) error {
	if g.player1.GetTeam() == p2.GetTeam() {
		return ErrTeamAlreadySelected
	}

	if g.state == GameStarted {
		return ErrCantJoinStartedGame
	}

	g.player2 = p2
	g.state = GameStarted

	return nil
}

func (g *Game) Play(p Player, s Statement) (Player, error) {
	if g.state == GameCreated {
		return nil, ErrWaitingPlayerJoin
	}

	if g.state == GameOver {
		return nil, ErrGameOverAlready
	}

	if g.lastRound != nil &&
		g.lastRound.player == p {

		return nil, ErrCanPlayTwice
	}

	if g.rounds[s] != nil {
		return nil, ErrInvalidStatement
	}

	r := &round{
		player:    p,
		statement: s,
	}

	g.rounds[s] = r
	g.lastRound = r

	// we need at least 5 rounds to be able to have a winner
	// then we must check
	g.roundsLeft--
	if g.roundsLeft < 6 {
		if g.hasWinner() {
			g.state = GameOver
			return p, nil
		}
		// tied
		if g.roundsLeft == 0 {
			g.state = GameOver
			return nil, nil
		}
	}

	return nil, nil
}

func (g *Game) hasWinner() bool {
	//rows
	if g.allSameTeam(Statement0, Statement1, Statement2) {
		return true
	}
	if g.allSameTeam(Statement3, Statement4, Statement5) {
		return true
	}
	if g.allSameTeam(Statement6, Statement7, Statement8) {
		return true
	}
	//cols
	if g.allSameTeam(Statement0, Statement3, Statement6) {
		return true
	}
	if g.allSameTeam(Statement1, Statement4, Statement7) {
		return true
	}
	if g.allSameTeam(Statement2, Statement5, Statement8) {
		return true
	}
	//diagonal
	if g.allSameTeam(Statement0, Statement4, Statement8) {
		return true
	}
	if g.allSameTeam(Statement2, Statement4, Statement6) {
		return true
	}

	return false
}

func (g *Game) allSameTeam(s1, s2, s3 Statement) bool {
	if g.rounds[s1] == nil || g.rounds[s2] == nil || g.rounds[s3] == nil {
		return false
	}

	return g.rounds[s1].player.GetTeam() == g.rounds[s2].player.GetTeam() && g.rounds[s2].player.GetTeam() == g.rounds[s3].player.GetTeam()
}
