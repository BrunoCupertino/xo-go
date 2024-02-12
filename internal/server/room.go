package server

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/BrunoCupertino/xo-go/internal/state"
)

type Room struct {
	acceptor ConnectionAcceptor
	game     *state.Game

	p1Conn net.Conn
	p2Conn net.Conn

	startChann chan struct{}
	p1Chann    chan []byte
	p2Chann    chan []byte
}

var errUnknownMessage = errors.New("unkown message")

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

		p1 := state.NewHumanPlayer(state.OTeam)

		r.game = state.NewGame(p1)
		r.p1Conn = playerConn
		r.p1Chann = make(chan []byte, 10)

		r.send(p1, "TO1")

		go readMsgsFromConnection(r.p1Conn, r.p1Chann)

		return
	}

	p2 := state.NewHumanPlayer(state.XTeam)

	r.game.Join(p2)
	r.p2Conn = playerConn
	r.p2Chann = make(chan []byte, 10)

	r.send(p2, "TX2")

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
			r.tryPlay(r.game.player1, string(p1Msg))
		case p2Msg := <-r.p2Chann:
			fmt.Printf("player 2 send this message: %s \n", string(p2Msg))
			r.tryPlay(r.game.player2, string(p2Msg))
		}
	}
}

func (r *Room) tryPlay(p state.Player, cmd string) error {
	if string(cmd[0]) != "S" {
		fmt.Printf("error while playing cmd %s", string(cmd[0]))
		return errUnknownMessage
	}

	if string(cmd[1]) != string(p.GetTeam()) {
		fmt.Printf("error while playing diff team %s", string(cmd[1]))
		return errUnknownMessage
	}

	s, _ := strconv.Atoi(string(cmd[2]))

	var statement Square = Square(byte(s))

	winner, err := r.game.Play(p, statement)
	if err != nil {
		fmt.Printf("error while playing %d", err)
		return err
	}

	fmt.Print("well played")

	if winner == nil {
		return r.broadcast(fmt.Sprintf("S%s%d", p.GetTeam(), statement))
	}

	fmt.Printf("game over, winner %s", winner.GetTeam())

	return r.broadcast(fmt.Sprintf("W%s%d", p.GetTeam(), statement))
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

func (r *Room) send(player state.Player, data string) error {
	if r.game.player1 == player {
		_, err := r.p1Conn.Write([]byte(data))
		return err
	}

	_, err := r.p2Conn.Write([]byte(data))
	return err
}

func (r *Room) broadcast(data string) error {
	r.send(r.game.player1, data)
	return r.send(r.game.player2, data)
}

// Room:
// Join(player): create or join the game (player, error)
// Play(Statement): play round (Statement, game status, error)

// ServerGameManager:
// New(): create room and async wait for players | async call onConnected(player)
// onConnected(conn): assign player to its connection and start listened msgs from conn
// Start(): consume msgs from players, decode and call room.Play()

// Statement:
// (team, square, game status)

// Encoder:
// DefaultEncoder
// Encode(Statement): encode statement into string then []byte
// Decode([]byte): decode data into string, then Statement

// Room, GameState must be private
