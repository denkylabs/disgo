package disgo

// Create a map, where keys are strings and values are functions with any type
var eventHandlers = make(map[string]interface{})

// Add a Discord event handler
// 	func ReadyEvent(data map[string]interface{}) {
//  	fmt.Println("Bot is ready!")
// 	}
// 	session.AddListener("READY", ReadyEvent)
func (s *Session) AddListener(eventName string, handler func(data map[string]interface{})) {
	eventHandlers[eventName] = handler
}

// Remove a Discord event handler
// 	session.RemoveListener("READY")
func (s *Session) RemoveListener(eventName string) {
	delete(eventHandlers, eventName)
}

// Manually trigger a Discord event
func (s *Session) triggerEvent(eventName string, eventData map[string]interface{}) {
	if eventHandlers[eventName] != nil {
		eventHandlers[eventName].(func(data map[string]interface{}))(eventData)
	}
}
