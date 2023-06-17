package main

import (
	"fmt"
)

// assumed number of validators (can be dynamic but have cosnt for simplicity)
const NUM_OF_VALIDATORS = 10

// struct describing duty corresponding to json received from websocket
type Request struct {
	Validator string `json:"validator"`
	Duty      string `json:"duty"`
	Height    int    `json:"height"`
}

sim

func main() {
	// create a list of validators
	validators := make([]*Validator, 0)

	// for waiting validator routines
	var wg sync.WaitGroup

	// run each validator duty processor separately
	for v := 0; v < NUM_OF_VALIDATORS; v++ {
		// create a new validator
		validator := &Validator{
			ValidatorID: v,
			// give a space (10) for receiving multiple requests not to block the main thread
			requests : make(chan Duty, 10)
		}

		// save validator
		validators = append(validators, validator)

		// start listening for incoming requests
		go run validator.ListenForRequests(&wg)
	}

	for DutyRequest :=




	// Couldn't get response during the assignment weather I should implement graceful shutdown or wait group or what. So for simplicity I wait by sleep
	time.Sleep(time.Duration * 5)
	return
}
