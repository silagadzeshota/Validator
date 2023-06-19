package wslistener

import (
  "log"
  "time"
  // just to note package is no longer maintained but used for simplicity
  "github.com/gorilla/websocket"
)

// Delay before attempting reconnection
const reconnectDelay = 5 * time.Second

// reads incoming data from websocket as a byte array and pushes into the output channel
// the implementation doesn't know about the "duties", separating websocket functionality from the logic.
// Also not implementing wss secure connection functionality for simplicity
func PullData(endpoint string, output chan []byte) {
	// Keep retrying the WebSocket connection
	for {
		// Attempt to establish a WebSocket connection
		conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
		if err != nil {
			log.Println("Failed to connect to WebSocket:", err)
			time.Sleep(reconnectDelay)
			continue // Retry the connection
		}

		// Connection successful, handle incoming messages
		handleWebSocket(conn, output)
	}
}

// read incoming data and push into chan
func handleWebSocket(conn *websocket.Conn, output chan []byte) {
	defer conn.Close()
	for {
		// Read a message from the WebSocket connection
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}

		// Process the received message
		output <- msg
	}
}
