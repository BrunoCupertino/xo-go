package server

import (
	"github.com/BrunoCupertino/xo-go/internal/state"
)

type Room struct {
	gameState *state.GameState
	winner    state.Player
}

func NewRoom() *Room {
	r := &Room{}

	return r
}

func (r *Room) CreatePlayerAndJoin() (state.Player, error) {
	if r.gameState == nil {
		p1 := state.NewHumanPlayer(state.OTeam)

		r.gameState = state.NewGameState(p1)

		return p1, nil
	}

	p2 := state.NewHumanPlayer(state.XTeam)

	return p2, r.gameState.Join(p2)
}

func (r *Room) Play(s state.Statement) (*state.Statement, state.GameStatus, error) {
	player := r.GetPlayerByTeam(s.Team)

	winner, err := r.gameState.Play(player, s.Square)
	if err != nil {
		return nil, r.gameState.GetStatus(), err
	}

	statement := s
	gameStatus := r.gameState.GetStatus()

	if gameStatus == state.GameOver {
		teamWinner := state.NoTeam

		if winner != nil {
			teamWinner = winner.GetTeam()
			r.winner = winner
		}

		statement = *state.NewStatement(state.GameOvered, teamWinner, s.Square)
	}

	return &statement, r.gameState.GetStatus(), err
}

func (r *Room) GetPlayerByTeam(t state.Team) state.Player {
	player := r.gameState.GetPlayer1()

	if player.GetTeam() == t {
		return player
	}

	return r.gameState.GetPlayer2()
}
