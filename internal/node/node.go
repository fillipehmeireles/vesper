package node

import "time"

type Node struct {
	Id int
	Name string
	IpAddr string
	LastSeenTS time.Time
}
