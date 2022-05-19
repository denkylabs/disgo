package disgo

import "github.com/gorilla/websocket"

type Session struct {
	// The Identify structure of this session
	Identify Identify `json:"identify"`
	// The gorilla websocket connection
	ws *websocket.Conn

	// The heartbeat interval that Discord sent us
	heartbeatInterval float64
	// Whether if the last sent heartbeat was acked by Discord
	heartbeatAcked bool
	// The timestamp of the last heartbeat
	lastHeartbeatAck float64
	// Session ID set after the READY event. Used for Resuming
	sessionId string
	// The current session sequence number. Used for Resuming
	sequence float64
}
