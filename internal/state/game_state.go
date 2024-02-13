package state

import (
	"errors"

	"github.com/google/uuid"
)

type GameStatus byte

const (
	GameCreated GameStatus = iota
	GameStarted
	GameOver
)

// Square represent tic tac toe squares as following:
// 0 | 1 | 2
// ---------
// 3 | 4 | 5
// ---------
// 6 | 7 | 8
type Square byte

const (
	Square0 Square = iota
	Square1
	Square2
	Square3
	Square4 // middle
	Square5
	Square6
	Square7
	Square8
)

var ErrCanPlayTwice = errors.New("can not play twice wait your turn")
var ErrTeamAlreadySelected = errors.New("team already selected")
var ErrCantJoinStartedGame = errors.New("can not join a started game")
var ErrWaitingPlayerJoin = errors.New("waiting player join the game")
var ErrGameOverAlready = errors.New("ow this game is over already, somebody got finished!")
var ErrInvalidSquare = errors.New("square invalid my brother")

type round struct {
	player Player
	square Square
}

type GameState struct {
	id     uuid.UUID
	status GameStatus

	// store round history
	lastRound  *round
	rounds     [9]*round
	roundsLeft byte

	// players of the game
	player1 Player
	player2 Player
}

func NewGameState(p1 Player) *GameState {
	return &GameState{
		id:         uuid.New(),
		status:     GameCreated,
		player1:    p1,
		roundsLeft: 9,
	}
}

func (g *GameState) Join(p2 Player) error {
	if g.status == GameStarted {
		return ErrCantJoinStartedGame
	}

	if g.player1.GetTeam() == p2.GetTeam() {
		return ErrTeamAlreadySelected
	}

	g.player2 = p2
	g.status = GameStarted

	return nil
}

func (g *GameState) Play(p Player, s Square) (Player, error) {
	if g.status == GameCreated {
		return nil, ErrWaitingPlayerJoin
	}

	if g.status == GameOver {
		return nil, ErrGameOverAlready
	}

	if g.lastRound != nil &&
		g.lastRound.player == p {

		return nil, ErrCanPlayTwice
	}

	if g.rounds[s] != nil {
		return nil, ErrInvalidSquare
	}

	r := &round{
		player: p,
		square: s,
	}

	g.rounds[s] = r
	g.lastRound = r

	// we need at least 5 rounds to be able to have a winner
	// then we must check
	g.roundsLeft--
	if g.roundsLeft < 6 {
		if g.hasWinner() {
			g.status = GameOver
			return p, nil
		}
		// tied
		if g.roundsLeft == 0 {
			g.status = GameOver
			return nil, nil
		}
	}

	return nil, nil
}

func (g *GameState) GetPlayer1() Player {
	return g.player1
}

func (g *GameState) GetPlayer2() Player {
	return g.player2
}

func (g *GameState) GetStatus() GameStatus {
	return g.status
}

func (g *GameState) hasWinner() bool {
	// I know it's weird maybe I will refactor and use bit flags
	//rows
	if g.allSameTeam(Square0, Square1, Square2) {
		return true
	}
	if g.allSameTeam(Square3, Square4, Square5) {
		return true
	}
	if g.allSameTeam(Square6, Square7, Square8) {
		return true
	}
	//cols
	if g.allSameTeam(Square0, Square3, Square6) {
		return true
	}
	if g.allSameTeam(Square1, Square4, Square7) {
		return true
	}
	if g.allSameTeam(Square2, Square5, Square8) {
		return true
	}
	//diagonal
	if g.allSameTeam(Square0, Square4, Square8) {
		return true
	}
	if g.allSameTeam(Square2, Square4, Square6) {
		return true
	}

	return false
}

func (g *GameState) allSameTeam(s1, s2, s3 Square) bool {
	if g.rounds[s1] == nil || g.rounds[s2] == nil || g.rounds[s3] == nil {
		return false
	}

	return g.rounds[s1].player.GetTeam() == g.rounds[s2].player.GetTeam() &&
		g.rounds[s2].player.GetTeam() == g.rounds[s3].player.GetTeam()
}
