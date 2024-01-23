package ws_v2

import (
	"encoding/json"
	"fmt"
	"time"
)

type Function struct {
	Type    string          `json:"function"`
	Payload json.RawMessage `json:"payload"`
}

const (
	EventSendMessage = "send_message"
	EventNewMessage  = "new_message"
	EventTopTipper   = "event-top-tipper"
)

type FunctionHandler func(fn Function, c *Client) error

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
	Data    string `json:"data"`
}

type NewMessageEvent struct {
	SendMessageEvent
	Sent time.Time `json:"sent"`
}

func SendMessageHandler(fn Function, c *Client) error {

	var chatevent SendMessageEvent
	if err := json.Unmarshal(fn.Payload, &chatevent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	// Prepare an Outgoing Message to others
	var broadMessage NewMessageEvent

	broadMessage.Sent = time.Now()
	broadMessage.Message = chatevent.Message
	broadMessage.From = chatevent.From

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	// Place payload into an Event
	var outgoingEvent Function
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewMessage
	// Broadcast to all other Clients
	for client := range c.manager.clients {
		// Only send to clients inside the same chatroom
		if client.subscribeWsContext == c.subscribeWsContext {
			client.egress <- outgoingEvent
		}

	}
	return nil
}
