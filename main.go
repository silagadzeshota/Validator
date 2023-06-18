package main

import (
	"os"

	"dutyprocessor/validator"
	"dutyprocessor/dutyprocessor"
)

const (
	// websocket url env key to connect to receive duties
	DutyWebsocketURL = "DUTY_WEBSOCKET_URL"
)

// for simplicity I do not use graceful shutdown
func main() {
	// create main duty orchestrator
	var processor dutyprocessor.DutyProcessor
	processor.Duties = make(chan []byte)
	processor.Validators = make(map[string]*validator.Validator)

	// start listening to websocket to push incoming duties into duty channel
	processor.StartWSListenerForDuties(os.Getenv(DutyWebsocketURL))

	// starts listening to incoming duties from duty channel and distributing them among validators to process
	processor.StartDutyProcessor()

	return
}
