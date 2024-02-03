package client

type Connector interface {
	Connect()
}

type TCPConnector struct {
	opts *ConnectorOpts
}

type ConnectorOpts struct {
	Addr      string
	OnMessage func(any)
}

func NewTCPConnector(opts *ConnectorOpts) *TCPConnector {
	return &TCPConnector{
		opts: opts,
	}
}

func (c *TCPConnector) Connect() {
}
