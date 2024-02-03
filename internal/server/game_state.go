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
	Statement0 = iota
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
	lastRound *round
	rounds    [9]*round

	// players of the game
	player1 Player
	player2 Player
}

func NewGame(p1 Player) *Game {
	return &Game{
		id:      uuid.New(),
		state:   GameCreated,
		player1: p1,
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

func (g *Game) Play(p Player, s Statement) error {
	if g.state == GameCreated {
		return ErrWaitingPlayerJoin
	}

	if g.state == GameOver {
		return ErrGameOverAlready
	}

	if g.lastRound != nil &&
		g.lastRound.player.GetTeam() == p.GetTeam() {

		return ErrCanPlayTwice
	}

	if g.rounds[s] != nil {
		return ErrInvalidStatement
	}

	r := &round{
		player:    p,
		statement: s,
	}

	g.rounds[s] = r
	g.lastRound = r

	// we need at least 5 rounds to be able to have a winner
	// then we must check
	if len(g.rounds) > 4 {
		//TODO check if has a winner and
	}

	return nil
}
