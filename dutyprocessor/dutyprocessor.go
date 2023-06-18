package dutyprocessor

import (
	"log"
	"encoding/json"

	"dutyprocessor/validator"
	"dutyprocessor/wslistener"
)

// main structure for orchestrating the duties
type DutyProcessor struct {
	// registered/active validator map
	Validators map[string]*validator.Validator

	// channel for incoming duties from websocket
	Duties chan []byte
}

// request json object incoming from websocket
type DutyRequest struct {
	Validator string `json:"validator"`
	Duty      string `json:"duty"`
	Height    int    `json:"height"`
}

/*
	Start listening to websocket and insert incoming data into channel.
	Websocket listener part is separated from duty validator as they are functionally and logically independent
*/
func (dutyProcessor *DutyProcessor) StartWSListenerForDuties(websocketURL string) {
	// start listening and pushing incoming data into channel asynchronously
	go wslistener.PullData(websocketURL, dutyProcessor.Duties)
}

func (dutyProcessor *DutyProcessor) StartDutyProcessor() {
	log.Println("listening for incoming duties to process")
  for {
    // wait for a new request
    duty := <- dutyProcessor.Duties

		// unmarshal duty
		var dutyRequest DutyRequest
		if err := json.Unmarshal([]byte(duty), &dutyRequest); err != nil {
			log.Println("Cannot unmarshal incoming data as a duty request")
			continue;
		}

		// if validator is not active create, register and start it
		if _, ok := dutyProcessor.Validators[dutyRequest.Validator]; ok == false {
			// create new validator
			dutyProcessor.CreateAndStartValidator(dutyRequest.Validator)
		}

		// push new duty to the validator asynchronously (in case it blocks)
		go dutyProcessor.Validators[dutyRequest.Validator].PushNewDuty(validator.Duty{Duty: dutyRequest.Duty, Height: dutyRequest.Height})
  }

  return
}

// creates validator, registers into dutyProcessor and starts validator to listen to incoming duties and process them
func (dutyProcessor *DutyProcessor) CreateAndStartValidator(validatorID string) {
	// create new validator
	newValidator := validator.NewValidator(validatorID)

	// register validator into duty processor
	dutyProcessor.Validators[validatorID] = newValidator

	//activate validator listening to incoming duties and processing them sequentially for heights
	go newValidator.Start()
}
