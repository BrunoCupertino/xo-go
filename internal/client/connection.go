package client

import (
	"errors"
	"fmt"
	"net"
)

type Connector interface {
	Connect() error
	Start()
	Send(data []byte) error
}

type TCPConnector struct {
	opts *TCPConnectorOpts
	conn net.Conn

	startChann     chan struct{}
	onMessageChann chan []byte
}

type TCPConnectorOpts struct {
	addr string
	// OnMessage func(any)
}

func NewTCPConnectorOpts(addr string) *TCPConnectorOpts {
	return &TCPConnectorOpts{
		addr: addr,
	}
}

func NewTCPConnector(opts *TCPConnectorOpts) *TCPConnector {
	return &TCPConnector{
		opts: opts,
		// startChann:     make(chan struct{}),
		onMessageChann: make(chan []byte, 10),
	}
}

func (c *TCPConnector) Connect() error {
	conn, err := net.Dial("tcp", c.opts.addr)
	if err != nil {
		return err
	}

	c.conn = conn

	go c.readMessages()

	fmt.Println("connected")

	//c.startChann <- struct{}{}

	fmt.Println("its about to start")

	return nil
}

func (c *TCPConnector) Start() {
	fmt.Println("waiting start2")

	// <-c.startChann

	fmt.Println("waiting done")

	for {

		select {
		case msg := <-c.onMessageChann:
			fmt.Printf("receive this message: %s \n", string(msg))
		}
	}
}

func (c *TCPConnector) Send(data []byte) error {
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

func (c *TCPConnector) readMessages() {
	buf := make([]byte, 1024)

	for {
		n, err := c.conn.Read(buf)
		if err != nil {
			fmt.Printf("stopping reading messages due to %s \n", err.Error())
			break
		}

		msg := make([]byte, n)
		copy(msg, buf[:n])

		c.onMessageChann <- msg
	}
}
