package helly

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Wait Group for sendAsyncMessage()
var messagesWaitGroup sync.WaitGroup

// Wait Group for handleWSMessages()
var incomingWg sync.WaitGroup

// Logs the client in, establishing a WebSocket connection to Discord
// This function blocks until the connection is lost, ended or an error occurs
func (s *Session) Connect() (e error) {
	DiscordWSUrl, err := getDiscordGatewayURL()

	if err != nil {
		return err
	}

	ws, _, err := websocket.DefaultDialer.Dial(DiscordWSUrl, nil)

	if err != nil {
		return err
	}

	s.ws = ws

	err = handleWSMessages(s)

	if err != nil {
		return err
	}

	return nil
}

// Function to parse incoming packets from Discord
func parsePacket(s *Session, packet map[string]interface{}) error {
	var packetEventName string
	var packetOpcode float64
	var packetData map[string]interface{}
	var packetSequence float64

	if packet["t"] != nil {
		packetEventName = packet["t"].(string)
	}

	if packet["op"] != nil {
		packetOpcode = packet["op"].(float64)
	}

	// Check if the packet has a valid data field
	if packet["d"] != nil && reflect.TypeOf(packet["d"]).Kind() != reflect.Bool {
		packetData = packet["d"].(map[string]interface{})
	}

	// Sequence
	if packet["s"] != nil {
		packetSequence = packet["s"].(float64)
		s.seq = packetSequence
	}

	fmt.Printf("Event name: %v\nEvent opcode: %v\nEvent data: %v\nEvent sequence: %v\n", packetEventName, packetOpcode, packetData, packetSequence)

	switch packetOpcode {
	// Debug Opcode
	case 0:
		parseEvent(packetEventName, packetData, s)

	// Hello Opcode
	case 10:
		s.heartbeatInterval = packetData["heartbeat_interval"].(float64)
		s.heartbeatAcked = true
		s.lastHeartbeatAck = float64(time.Now().UnixMilli())
		fmt.Println("Defined heartbeat interval to", s.heartbeatInterval)

		sendHeartbeat(s)
		sendIdentify(s)

	// Invalid Session Opcode
	case 9:
		return errors.New("invalid session")

	// Heartbeat Opcode
	case 11:
		s.heartbeatAcked = true
		s.lastHeartbeatAck = float64(time.Now().UnixMilli())
	}

	return nil
}

// Function to handle websocket messages
// The ReadMessage() function blocks until a message is received, thats why we need to use a goroutine
func handleWSMessages(s *Session) (e error) {
	var finalErr string

	incomingWg.Add(1)
	go func() {
		for {
			_, message, err := s.ws.ReadMessage()

			if err != nil {
				finalErr = fmt.Sprint(err)
				break
			}

			strMessage := string(message)
			// Parse the string to a JSON object and save it to a variable, so parseEvent can use it
			var jsonMessage map[string]interface{}
			err = json.Unmarshal([]byte(strMessage), &jsonMessage)

			if err != nil {
				finalErr = fmt.Sprint(err)
				break
			}

			fmt.Println(strMessage)
			parsePacket(s, jsonMessage)
		}

		// When the loop is broken, close the websocket
		incomingWg.Done()
		s.ws.Close()
	}()

	// Wait until the websocket is closed or an error occurs
	incomingWg.Wait()

	if finalErr != "" {
		return errors.New(finalErr)
	}

	return nil
}

/* Websocket send functions */
func sendIdentify(s *Session) {
	sendAsyncMessage(s, GatewayIdentifySendPacket{2, s.Identify})
}

func sendHeartbeat(s *Session) {
	sendAsyncMessage(s, GatewayHeartbeatSendPacket{1, 0})
}

func sendAsyncMessage(s *Session, message interface{}) error {
	var err error
	messagesWaitGroup.Add(1)
	go func() {
		defer messagesWaitGroup.Done()
		err = s.ws.WriteJSON(message)
	}()
	messagesWaitGroup.Wait()

	return err
}

/* Function to parse dispatch events */
func parseEvent(eventName string, eventData map[string]interface{}, s *Session) {
	switch eventName {
	case "READY":
		triggerReadyDispatchEvent(eventData, s)
	}
}

// Gateway READY dispatch event
func triggerReadyDispatchEvent(eventData map[string]interface{}, s *Session) {
	s.sId = eventData["session_id"].(string)
}
