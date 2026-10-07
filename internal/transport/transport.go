package transport


type Peer struct {
    Address string
	Port int
}

type Transport interface {
    Send(msg []byte, address string) error
	Receive() ([]byte, Peer, error)
	CloseConnection () error
}

type Broadcaster interface {
    Broadcast(payload []byte) error
}
