package server

import (
	"fmt"
	"net"
)

type Room struct {
	acceptor ConnectionAcceptor
	game     *Game

	p1Conn net.Conn
	p2Conn net.Conn
}

func NewRoom(acceptor ConnectionAcceptor) *Room {
	r := &Room{
		acceptor: acceptor,
	}

	go r.waitingPlayers()

	return r
}

func (r *Room) onConnected(playerConn net.Conn) {
	if r.game == nil {

		p1 := NewHumanPlayer(OTeam)
		r.game = NewGame(p1)
		r.p1Conn = playerConn

		return
	}

	p2 := NewHumanPlayer(XTeam)

	r.game.Join(p2)
	r.p2Conn = playerConn
}

func (r *Room) waitingPlayers() {
	defer r.acceptor.Close()

	for i := 0; i < 2; i++ {
		fmt.Println("waiting players joining room...")

		p1Conn, _ := r.acceptor.Listen()

		fmt.Println("player 1 has joinned the room")

		r.onConnected(p1Conn)

		fmt.Println("waiting player 2 joining room...")

		p2Conn, _ := r.acceptor.Listen()

		fmt.Println("player 2 has joinned the room")

		r.onConnected(p2Conn)
	}
}
