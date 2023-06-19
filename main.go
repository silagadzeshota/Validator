package main

import (
	"os"

	"dutyprocessor/dutyprocessor"
)

const (
	// websocket url env key to connect to receive duties
	DutyWebsocketURL = "DUTY_WEBSOCKET_URL"
)

// for simplicity I do not use graceful shutdown
func main() {
	// create main duty orchestrator
	processor := dutyprocessor.NewDutyProcessor()

	// start listening to websocket to push incoming duties into duty channel
	processor.StartWSListenerForDuties(os.Getenv(DutyWebsocketURL))

	// starts listening to incoming duties from duty channel and distributing them among validators to process
	processor.StartDutyProcessor()

	return
}
