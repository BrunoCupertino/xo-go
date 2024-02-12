package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	var (
		p1 = NewHumanPlayer(XTeam)
		g  = NewGame(p1)
	)

	assert.Equal(t, GameCreated, g.state)
	assert.Equal(t, p1, g.player1)
}

func TestJoin(t *testing.T) {
	var (
		p1 = NewHumanPlayer(XTeam)
		p2 = NewHumanPlayer(XTeam)
		g  = NewGame(p1)
	)

	err := g.Join(p2)

	assert.ErrorIs(t, err, ErrTeamAlreadySelected)
	assert.Equal(t, GameCreated, g.state)
	assert.Nil(t, g.player2)

	p2 = NewHumanPlayer(OTeam)

	err = g.Join(p2)

	assert.NoError(t, err)
	assert.Equal(t, GameStarted, g.state)
	assert.Equal(t, p2, g.player2)

	err = g.Join(p2)
	assert.ErrorIs(t, err, ErrCantJoinStartedGame)
	assert.Equal(t, GameStarted, g.state)
}

func TestPlayValidations(t *testing.T) {
	var (
		p1 = NewHumanPlayer(XTeam)
		p2 = NewHumanPlayer(OTeam)
		g  = NewGame(p1)
	)

	winner, err := g.Play(p1, Square0)
	assert.ErrorIs(t, err, ErrWaitingPlayerJoin)
	assert.Nil(t, winner)

	g.Join(p2)

	winner, err = g.Play(p1, Square0)
	assert.NoError(t, err)
	assert.Nil(t, winner)

	winner, err = g.Play(p1, Square0)
	assert.ErrorIs(t, err, ErrCanPlayTwice)
	assert.Nil(t, winner)

	winner, err = g.Play(p2, Square0)
	assert.ErrorIs(t, err, ErrInvalidSquare)

	winner, err = g.Play(p2, Square1)
	assert.NoError(t, err)
	assert.Nil(t, winner)
}

func TestPlayWinner(t *testing.T) {
	var (
		p1 = NewHumanPlayer(XTeam)
		p2 = NewHumanPlayer(OTeam)
		g  = NewGame(p1)
	)

	g.Join(p2)

	winner, _ := g.Play(p1, Square0)
	assert.Nil(t, winner)

	winner, _ = g.Play(p2, Square3)
	assert.Nil(t, winner)

	winner, _ = g.Play(p1, Square1)
	assert.Nil(t, winner)

	winner, _ = g.Play(p2, Square4)
	assert.Nil(t, winner)

	winner, _ = g.Play(p1, Square2)
	assert.NotNil(t, winner)
	assert.Equal(t, p1.GetTeam(), winner.GetTeam())
}
