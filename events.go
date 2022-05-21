package disgo

// Create a map, where keys are strings and values are functions with any type
var eventHandlers = make(map[string]interface{})

// Add a Discord event handler
// 	func ReadyEvent(data map[string]interface{}) {
//  	fmt.Println("Bot is ready!")
// 	}
// 	session.AddListener("READY", ReadyEvent)
func (*Session) AddListener(eventName string, handler func(data map[string]interface{})) {
	eventHandlers[eventName] = handler
}

// Remove a Discord event handler
// 	session.RemoveListener("READY")
func (*Session) RemoveListener(eventName string) {
	delete(eventHandlers, eventName)
}

// Manually trigger a Discord event
func (*Session) triggerEvent(eventName string, eventData map[string]interface{}) {
	if eventHandlers[eventName] != nil {
		eventHandlers[eventName].(func(data map[string]interface{}))(eventData)
	}
}

/* Discord events structure */

// Represents the Discord HELLO event
type APIGatewayHelloDiscordDispatchEvent struct {
	// The opcode of this event
	Op int `json:"op"`
	// The name of the received event
	T int `json:"t"`
	// The sequence number of the last event received by the client
	S int `json:"s"`
	// The data of the event
	Data struct {
		// the interval (in milliseconds) the client should heartbeat with
		HeartbeatInterval int `json:"heartbeat_interval"`
	} `json:"d"`
}

// Represents the Discord READY event
type APIGatewayReadyDiscordDispatchEvent struct {
	// The opcode of this event
	Op int `json:"op"`
	// The name of the received event
	T int `json:"t"`
	// The sequence number of the last event received by the client
	S int `json:"s"`
	// The data of the event
	Data struct {
		// gateway version
		Version int `json:"v"`
		// information about the user incluiding email
		User APIUser `json:"user"`
		// the guilds the user is in (slice of Unavailable Guild objects)
		Guilds []APIUnavailableGuild `json:"guilds"`
		// the session id used for resuming
		SessionId string `json:"session_id"`
		// the shard information associated with this session, if sent when identifying
		Shard []int `json:"_shards"`
	} `json:"d"`
}

// Represents the Discord RESUMED event
type APIGatewayResumedDiscordDispatchEvent struct {
	// The opcode of this event
	Op int `json:"op"`
	// The name of the received event
	T int `json:"t"`
	// The sequence number of the last event received by the client
	S int `json:"s"`
}

// Represents the Discord RECONNECT event
type APIGatewayReconnectDiscordDispatchEvent struct {
	// The opcode of this event
	Op int `json:"op"`
	// The name of the received event
	T int `json:"t"`
	// The sequence number of the last event received by the client
	S int `json:"s"`
	// The data of the event (always null)
	Data interface{} `json:"d"`
}