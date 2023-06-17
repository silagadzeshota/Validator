package main

import (
	"time"
	"strconv"

	"validator/validator"
)

// assumed number of validators (can be dynamic but have cosnt for simplicity)
const NUM_OF_VALIDATORS = 10

// struct describing duty corresponding to json received from websocket
type Request struct {
	Validator string `json:"validator"`
	Duty      string `json:"duty"`
	Height    int    `json:"height"`
}

// instead of using websocket I'm declaring some static request array to simulate the process.
// Websocket endpoint couldn't work with existing golang libraries as they strictly define url scheme (must start with ws or wss)
// I could change library code (not good) or rewrite python (questionably). As I couldn't know what to do (as I couldn't receive any answer about this)
// I just declare the array to simulate the process without python process submitting requests through websocket
var requests = []Request{
	Request {
			Validator: "2",
			Duty: "ATTESTER",
			Height: 8,
	},
	Request {
			Validator: "9",
			Duty: "SYNC_COMMITTEE",
			Height: 2,
	},
	Request {
			Validator: "3",
			Duty: "PROPOSER",
			Height: 5,
	},
	Request {
			Validator: "4",
			Duty: "AGGREGATOR",
			Height: 4,
	},
	Request {
			Validator: "6",
			Duty: "ATTESTER",
			Height: 21,
	},
}

func main() {
	// create a list of validators
	validators := make([]*validator.Validator, 0)

	// run each validator duty processor separately
	for v := 0; v < NUM_OF_VALIDATORS; v++ {
		// create a new validator
		validator := &validator.Validator{
			ValidatorID: v,
			// give a space (10) for receiving multiple requests not to block the main thread
			Requests : make(chan validator.Duty, 10),
		}

		// save validator
		validators = append(validators, validator)

		// start listening for incoming requests
		go validator.ListenForRequests()
	}

	// simulate receiving requests from websocket and processing them
	for _, dutyRequest := range requests {
		// assuming fields are parsed and validated
		validatorID, _ := strconv.Atoi(dutyRequest.Validator)
		validators[validatorID].Requests <- validator.Duty{Duty: dutyRequest.Duty, Height: dutyRequest.Height}
	}

	// Couldn't get response during the assignment weather I should implement graceful shutdown or wait group or what. So for simplicity I wait by sleeping
	time.Sleep(time.Second * 3)
	return
}
