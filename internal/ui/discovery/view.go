package discovery

import (
	"fmt"
	"strconv"
	"time"

	"charm.land/bubbles/v2/table"
)



func (m Model) Render() string {
	var peerTable []table.Row
	for _,peer := range <- m.peerTableChan {
		fmt.Println(peer)
		peerTable = append(peerTable, table.Row{
			peer.Name,
			peer.IpAddr,
			strconv.FormatFloat(time.Since(peer.LastSeenTS).Seconds(),'f',2,64),
		})
	}

	m.peerTable.SetRows(peerTable)
	return m.peerTable.View()
}
