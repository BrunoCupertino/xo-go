package server

import (
	"fmt"
	"net"
)

type Room struct {
	acceptor ConnectionAcceptor
	game     Game

	p1Conn net.Conn
	p2Conn net.Conn
}

func (r *Room) NewRoom(acceptor ConnectionAcceptor) *Room {
	room := &Room{}

	go waitingPlayers(room)

	return room
}

func (r *Room) Join(p2 Player) error {
	return r.game.Join(p2)
}

// replace for onAccepted event
func waitingPlayers(room *Room) {
	for i := 0; i < 2; i++ {
		fmt.Println("waiting players joining room...")

		p1Conn, _ := room.acceptor.Listen()

		p1 := NewHumanPlayer(OTeam)

		room.game = *NewGame(p1)
		room.p1Conn = p1Conn

		fmt.Println("waiting player 2 joining room...")

		p2Conn, _ := room.acceptor.Listen()

		p2 := NewHumanPlayer(XTeam)

		room.game.Join(p2)
		room.p2Conn = p2Conn
	}
}
