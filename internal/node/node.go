package node

import (
	"time"
)

type Node struct {
	ID string
	Name string
	IpAddr string
	Port int
	LastSeenTS time.Time
}
