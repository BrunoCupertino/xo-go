package client

import "github.com/BrunoCupertino/xo-go/internal/state"

type ClientGameState struct {
	round    byte
	myTeam   state.Team
	lastTeam state.Team
	winner   state.Team
	change   state.StateChange
	board    [9]state.Team
}

func NewClientGameState(t state.Team) *ClientGameState {
	s := &ClientGameState{
		myTeam: t,
	}

	for i := 0; i < len(s.board); i++ {
		s.board[i] = " "
	}

	return s
}

func (s *ClientGameState) Change(statement *state.Statement) {
	s.change = statement.State

	if statement.State == state.TeamSelected {
		return
	}

	s.round++
	s.board[statement.Square] = statement.Team
	s.lastTeam = statement.Team

	if statement.State == state.GameOvered {
		s.winner = statement.Team
	}
}

func (s *ClientGameState) GetBoard() (byte, [9]state.Team) {
	return s.round, s.board
}

func (s *ClientGameState) MyTeam() state.Team {
	return s.myTeam
}

func (s *ClientGameState) CurrentChange() state.StateChange {
	return s.change
}

func (s *ClientGameState) Winner() state.Team {
	return s.winner
}

func (s *ClientGameState) IsMyTurn() bool {
	return s.myTeam != s.lastTeam
}
