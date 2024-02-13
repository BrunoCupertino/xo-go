package server

import (
	"fmt"
	"net"

	"github.com/BrunoCupertino/xo-go/internal/encoding"
	"github.com/BrunoCupertino/xo-go/internal/state"
)

type byPlayer struct {
	player  state.Player
	conn    net.Conn
	msgChan chan []byte
}

func newByPlayer(p state.Player, c net.Conn) *byPlayer {
	bp := &byPlayer{
		player:  p,
		conn:    c,
		msgChan: make(chan []byte),
	}

	go readMessages(bp.conn, bp.msgChan)

	return bp
}

type GameManager struct {
	acceptor  ConnectionAcceptor
	encoder   encoding.StatementEncoder
	room      *Room
	startChan chan struct{}
	stopChan  chan struct{}
	p1        *byPlayer
	p2        *byPlayer
}

func NewGameManager(a ConnectionAcceptor, e encoding.StatementEncoder) *GameManager {
	m := &GameManager{
		acceptor:  a,
		encoder:   e,
		startChan: make(chan struct{}),
		stopChan:  make(chan struct{}),
	}

	return m
}

func (m *GameManager) Start() {
	m.room = NewRoom()

	go m.waitingForPlayers()

	<-m.startChan

	for {
		select {
		case msg := <-m.p1.msgChan:
			m.process(m.p1, msg)
		case msg := <-m.p2.msgChan:
			m.process(m.p1, msg)
		}
	}
}

func (m *GameManager) Stop() {
	<-m.stopChan
}

func (m *GameManager) process(bp *byPlayer, msg []byte) error {
	statement, err := m.encoder.Decode(msg)
	if err != nil {
		return err
	}

	statement, _, err = m.room.Play(*statement)
	if err != nil {
		return err
	}

	return m.broadcast(statement)
}

func (m *GameManager) onConnected(playerConn net.Conn) {
	player, err := m.room.CreatePlayerAndJoin()
	if err != nil {
		fmt.Printf("error while joining room %s", err.Error())

		return
	}

	bp := newByPlayer(player, playerConn)

	m.send(bp, state.NewStatement(state.TeamSelected, player.GetTeam(), 0))

	if m.p1 == nil {
		m.p1 = bp
		return
	}

	m.p2 = bp

	m.startChan <- struct{}{}
}

func (m *GameManager) send(bp *byPlayer, statement *state.Statement) error {
	data, _ := m.encoder.Encode(statement)

	_, err := bp.conn.Write([]byte(data))

	return err
}

func (m *GameManager) broadcast(statement *state.Statement) error {
	_ = m.send(m.p1, statement)
	err := m.send(m.p2, statement)

	return err
}

func (m *GameManager) waitingForPlayers() {
	defer m.acceptor.Close()

	fmt.Println("waiting players joining room...")

	for i := 0; i < 2; i++ {

		playerConn, _ := m.acceptor.ListenAndAccept()

		m.onConnected(playerConn)
	}
}

func readMessages(conn net.Conn, onMessageChann chan<- []byte) {
	buf := make([]byte, 10)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("could not read message due to %s", err.Error())
			break
		}

		msg := make([]byte, n)
		copy(msg, buf[:n])

		onMessageChann <- msg
	}
}
