package client

import (
	"errors"
	"fmt"
	"net"

	"github.com/BrunoCupertino/xo-go/internal/encoding"
	"github.com/BrunoCupertino/xo-go/internal/state"
)

type GameClientManager interface {
	Start()
	Stop()
	Send(s *state.Square) error
}

type TCPGameClientManager struct {
	opts       *TCPGameClientManagerOpts
	conn       net.Conn
	encoder    encoding.StatementEncoder
	state      *ClientGameState
	renderer   Renderer
	startChann chan struct{}
	msgChan    chan []byte
}

type TCPGameClientManagerOpts struct {
	addr string
}

func NewTCPGameClientManagerOpts(addr string) *TCPGameClientManagerOpts {
	return &TCPGameClientManagerOpts{
		addr: addr,
	}
}

func NewTCPGameClientManager(opts *TCPGameClientManagerOpts,
	e encoding.StatementEncoder,
	r Renderer) *TCPGameClientManager {

	return &TCPGameClientManager{
		opts:     opts,
		encoder:  e,
		renderer: r,
		msgChan:  make(chan []byte),
	}
}

func (c *TCPGameClientManager) connect() error {
	conn, err := net.Dial("tcp", c.opts.addr)
	if err != nil {
		return err
	}

	c.conn = conn

	fmt.Println("connected")

	return nil
}

func (c *TCPGameClientManager) Start() {
	err := c.connect()
	if err != nil {
		panic(err)
	}

	go c.readMessages()

	for {
		select {
		case msg := <-c.msgChan:
			c.process(msg)
		}
	}
}

func (c *TCPGameClientManager) Send(s state.Square) error {
	statement := state.NewBoardStatement(c.state.myTeam, s)

	data, err := c.encoder.Encode(statement)
	if err != nil {
		return err
	}

	n, err := c.conn.Write(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		fmt.Println("warning: seems we are losing data")
		return errors.New("data loss")
	}

	return nil
}

func (c *TCPGameClientManager) process(msg []byte) error {
	statement, err := c.encoder.Decode(msg)
	if err != nil {
		return err
	}

	if statement.State == state.TeamSelected {
		c.state = NewClientGameState(statement.Team)
	}

	c.state.Change(statement)
	c.renderer.Render(c.state)

	return nil
}

func (c *TCPGameClientManager) readMessages() {
	buf := make([]byte, 10)

	for {
		n, err := c.conn.Read(buf)
		if err != nil {
			fmt.Printf("stopping reading messages due to %s \n", err.Error())
			break
		}

		msg := make([]byte, n)
		copy(msg, buf[:n])

		c.msgChan <- msg
	}
}
