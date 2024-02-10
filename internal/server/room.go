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

		playerConn, _ := r.acceptor.Listen()

		fmt.Printf("player %d has joinned the room\n", i+1)

		r.onConnected(playerConn)
	}

	fmt.Println("waiting finished")
}
