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

	startChann chan struct{}
	p1Chann    chan []byte
	p2Chann    chan []byte
}

func NewRoom(acceptor ConnectionAcceptor) *Room {
	r := &Room{
		acceptor:   acceptor,
		startChann: make(chan struct{}),
	}

	go r.waitingPlayers()

	return r
}

func (r *Room) onConnected(playerConn net.Conn) {
	if r.game == nil {

		p1 := NewHumanPlayer(OTeam)

		r.game = NewGame(p1)
		r.p1Conn = playerConn
		r.p1Chann = make(chan []byte, 10)

		go readMsgsFromConnection(r.p1Conn, r.p1Chann)

		return
	}

	p2 := NewHumanPlayer(XTeam)

	r.game.Join(p2)
	r.p2Conn = playerConn
	r.p2Chann = make(chan []byte, 10)

	go readMsgsFromConnection(r.p2Conn, r.p2Chann)

	r.startChann <- struct{}{}
}

func (r *Room) Start() {
	fmt.Println("waiting start")

	<-r.startChann

	fmt.Println("waiting done")

	for {

		select {
		case p1Msg := <-r.p1Chann:
			fmt.Printf("player 1 send this message: %s \n", string(p1Msg))
		case p2Msg := <-r.p2Chann:
			fmt.Printf("player 2 send this message: %s \n", string(p2Msg))
		}
	}
}

func readMsgsFromConnection(conn net.Conn, onMessageChann chan<- []byte) {
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("could not read message due to %s", err)
			break
		}

		msg := make([]byte, n)
		copy(msg, buf[:n])

		onMessageChann <- msg
	}
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
