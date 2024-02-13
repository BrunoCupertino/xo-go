package server

import (
	"github.com/BrunoCupertino/xo-go/internal/state"
)

type Room struct {
	game   *state.Game
	winner state.Player
}

func NewRoom() *Room {
	r := &Room{}

	return r
}

func (r *Room) CreatePlayerAndJoin() (state.Player, error) {
	if r.game == nil {
		p1 := state.NewHumanPlayer(state.OTeam)

		r.game = state.NewGame(p1)

		return p1, nil
	}

	p2 := state.NewHumanPlayer(state.XTeam)

	return p2, r.game.Join(p2)
}

func (r *Room) Play(s state.Statement) (*state.Statement, state.GameState, error) {
	player := r.GetPlayerByTeam(s.Team)

	winner, err := r.game.Play(player, s.Square)
	if err != nil {
		return nil, r.game.GetStatus(), err
	}

	statement := s
	gameStatus := r.game.GetStatus()

	if gameStatus == state.GameOver {
		teamWinner := state.NoTeam

		if winner != nil {
			teamWinner = winner.GetTeam()
			r.winner = winner
		}

		statement = *state.NewStatement(state.GameOvered, teamWinner, s.Square)
	}

	return &statement, r.game.GetStatus(), err
}

func (r *Room) GetPlayerByTeam(t state.Team) state.Player {
	player := r.game.GetPlayer1()

	if player.GetTeam() == t {
		return player
	}

	return r.game.GetPlayer2()
}
