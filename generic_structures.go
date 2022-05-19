package disgo

import "github.com/gorilla/websocket"

type Session struct {
	// The Identify structure of this session
	Identify Identify `json:"identify"`
	// The gorilla websocket connection
	ws *websocket.Conn

	// The heartbeat interval that Discord sent us
	heartbeatInterval float64
	// If the last heartbeat was acked by Discord
	heartbeatAcked bool
	// Timestamp of the last heartbeat
	lastHeartbeatAck float64
	// Session ID set after the READY event
	sId string
	// The current session sequence number
	seq float64
}
