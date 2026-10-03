package node

import "time"

type Node struct {
	Name string
	IpAddr string
	LastSeenTS time.Time
}
