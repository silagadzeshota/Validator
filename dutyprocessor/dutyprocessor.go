package dutyprocessor

import (
	"log"
	"errors"
	"strconv"
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

// Create new DutyProcessor
func NewDutyProcessor() *DutyProcessor {
	// create and initialze dutyProcessor
	var processor DutyProcessor
	processor.Duties = make(chan []byte)
	processor.Validators = make(map[string]*validator.Validator)
	return &processor
}

/*
	Start listening to websocket and insert incoming data into channel.
	Websocket listener part is separated from duty validator as they are functionally and logically independent
*/
func (dutyProcessor *DutyProcessor) StartWSListenerForDuties(websocketURL string) {
	// start listening and pushing incoming data into channel asynchronously
	go wslistener.PullData(websocketURL, dutyProcessor.Duties)
}

// to start duty processor asynchronously for testing
func (dutyProcessor *DutyProcessor) StartDutyProcessorAsync() {
	go dutyProcessor.StartDutyProcessor()
}

// start waitinf for incoming duty requests and distribute them among validators
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

		// process request
		err := dutyProcessor.ProcessRequest(dutyRequest)
		if err != nil {
			log.Println(err.Error())
		}
  }
  return
}

// process pecific request. If corresponding validator doesn't exist create, register and start before submitting duty
func (dutyProcessor *DutyProcessor) ProcessRequest(dutyRequest DutyRequest) error {
	// validate duty request
	if err := dutyProcessor.validateDuty(dutyRequest); err != nil {
		return errors.New("Validation failed: " + err.Error())
	}

	// if validator is not active create, register and start it
	if _, ok := dutyProcessor.Validators[dutyRequest.Validator]; ok == false {
		// create new validator
		dutyProcessor.CreateAndStartValidator(dutyRequest.Validator)
	}

	// push new duty to the validator asynchronously (in case it blocks)
	go dutyProcessor.Validators[dutyRequest.Validator].PushNewDuty(validator.Duty{Duty: dutyRequest.Duty, Height: dutyRequest.Height})

	return nil
}

// validate duty
func (dutyProcessor *DutyProcessor) validateDuty(dutyRequest DutyRequest) error {
	// check duty type
	if dutyRequest.Duty != validator.Proposer && dutyRequest.Duty != validator.Attester && dutyRequest.Duty != validator.Aggregator && dutyRequest.Duty != validator.SyncCommittee {
		return errors.New("Duty type not recognized!")
	}

	// check height
	if dutyRequest.Height < 0 {
		return errors.New("Duty height not valid")
	}

	// check validator id to be integer
	if _, err := strconv.Atoi(dutyRequest.Validator); err != nil {
	    return errors.New("Validator ID not valid")
	}

	return nil
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

// returns validator corresponding to its ID
func (dutyProcessor *DutyProcessor) GetValidator(validatorID string) *validator.Validator {
	return dutyProcessor.Validators[validatorID]
}

// check if validator exists
func (dutyProcessor *DutyProcessor) ValidatorExists(validatorID string) bool {
	_, ok := dutyProcessor.Validators[validatorID]
	return ok
}
