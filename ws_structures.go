package disgo

// A structure representing the properties of a Identify object
type Properties struct {
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
	Os      string `json:"$os"`
}

// A structure representing a Discord session
type Identify struct {
	// Token used to start a new session
	Token          string     `json:"token"`
	Intents        int        `json:"intents"`
	ShardID        int        `json:"shard_id"`
	ShardCount     int        `json:"shard_count"`
	Compress       bool       `json:"compress"`
	LargeThreshold int        `json:"large_threshold"`
	Properties     Properties `json:"properties"`
}

// Packet structures
type GatewayHeartbeatSendPacket struct {
	Op   int   `json:"op"`
	Data int64 `json:"d"`
}

type GatewayIdentifySendPacket struct {
	Op   int      `json:"op"`
	Data Identify `json:"d"`
}
