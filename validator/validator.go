package validator

import (
  "log"
  "sync"
)

// describing specific duty for specific validator
type Duty struct {
	Duty      string
	Height    int
}

/*
status for each duty. As stated in the assignment we have only these types of duties and they never change,
that's the reason of hardcoding them. Otherwise those kind of types can be separated as a list.
Value explanation: for specific height duty:
0 means it's not received
1 means it's received but not processed
2 means it's received and being processed
3 means it's processed
*/
type DutyStatuses struct {
  Proposer      int
  Attester      int
  Aggregator    int
  SyncCommittee int
}

// type names
const (
  // Duty names
  Proposer      = "PROPOSER"
  Attester      = "ATTESTER"
  Aggregator    = "AGGREGATOR"
  SyncCommittee = "SYNC_COMMITTEE"

  // Duty status meanings
  DutyNotReceived      = 0
  DutyReceived         = 1
  DutyIsBeingProcessed = 2
  DutyProcessed        = 3
)

// define each validator
type Validator struct {
  CurrentHeight      int
  Duties             map[int]*DutyStatuses
  DutiesLocker       sync.Mutex
	ValidatorID        string
	Requests           chan Duty
  processingFinished chan int
}

// creates and returns pointer to new validator
func NewValidator(validatorID string) *Validator {
  // create and register validator
	validator := &Validator{
		CurrentHeight:      0,
		ValidatorID:        validatorID,
		Requests:           make(chan Duty),
		Duties:             make(map[int]*DutyStatuses),
    processingFinished: make(chan int),
	}

  return validator
}

// ProcessDuty receives duty and pushes into the channel for processing
func (v *Validator) PushNewDuty(duty Duty) {
  v.Requests <- duty
}

// Start starts validator which listens to incoming duties and processes them according to height
func (v *Validator) Start() {
  log.Println("Validator ", v.ValidatorID, " created and started listening for incoming duties ")

  for {
    // wait until we can receive value from one of the channels
    select {
    // receiving new duties
    case duty := <- v.Requests:
      // Process duty request just registers request for specific validator to process
      v.processDutyRequest(duty)
    case <- v.processingFinished:
      // meaning one of the duty processing finished
    }

    // after receiving value from any of the channels we check for the next work if it's available
    v.checkNextWork()
  }
  return
}

// Process duty request
func (v *Validator) processDutyRequest(duty Duty) {
  // lock map for safety
  v.DutiesLocker.Lock()
  defer v.DutiesLocker.Unlock()

  // check if height in the duty is activated (if we have anything received for specific height)
  if _, ok := v.Duties[duty.Height]; ok == false {
    v.Duties[duty.Height] = &DutyStatuses{}
  }

  // Use switch on duty type.
  switch {
  case duty.Duty == Proposer:
    if v.Duties[duty.Height].Proposer >= DutyReceived {
      log.Println("Validator ", v.ValidatorID, ": Duty ", Proposer, " already received")
      return
    } else {
      v.Duties[duty.Height].Proposer = 1
    }
  case duty.Duty == Aggregator:
    if v.Duties[duty.Height].Aggregator >= DutyReceived {
      log.Println("Validator ", v.ValidatorID, ": Duty ", Aggregator, " already received")
      return
    } else {
      v.Duties[duty.Height].Aggregator = 1
    }
  case duty.Duty == SyncCommittee:
    if v.Duties[duty.Height].SyncCommittee >= DutyReceived {
      log.Println("Validator ", v.ValidatorID, ": Duty ", SyncCommittee, " already received")
      return
    } else {
      v.Duties[duty.Height].SyncCommittee = 1
    }
  case duty.Duty == Attester:
    if v.Duties[duty.Height].Attester >= DutyReceived {
      log.Println("Validator ", v.ValidatorID, ": Duty ", Attester, " already received")
      return
    } else {
      v.Duties[duty.Height].Attester = 1
    }
  }

  // log receiving new duty
  log.Println("Validator ", v.ValidatorID, ": Received new duty ", duty.Duty, " for the height ", duty.Height)
  return
}

// for the current height check if there is a received duty we can process
func (v *Validator) checkNextWork() {
  // protect map
  v.DutiesLocker.Lock()
  defer v.DutiesLocker.Unlock()

  // check if height in the duty is activated (if we have anything received for specific height)
  if _, ok := v.Duties[v.CurrentHeight]; ok == false {
    v.Duties[v.CurrentHeight] = &DutyStatuses{}
  }

  // check if we have duty that we received and isn't
  if v.Duties[v.CurrentHeight].Proposer == DutyReceived {
    // mark duty as being processed and start processing it
    v.Duties[v.CurrentHeight].Proposer = DutyIsBeingProcessed
    go v.processDuty(v.CurrentHeight, Proposer)
  }

  // check if we have duty that we received and isn't
  if v.Duties[v.CurrentHeight].Attester == DutyReceived {
    // mark duty as being processed and start processing it
    v.Duties[v.CurrentHeight].Attester = DutyIsBeingProcessed
    go v.processDuty(v.CurrentHeight, Attester)
  }

  // check if we have duty that we received and isn't
  if v.Duties[v.CurrentHeight].Aggregator == DutyReceived {
    // mark duty as being processed and start processing it
    v.Duties[v.CurrentHeight].Aggregator = DutyIsBeingProcessed
    go v.processDuty(v.CurrentHeight, Aggregator)
  }

  // check if we have duty that we received and isn't
  if v.Duties[v.CurrentHeight].SyncCommittee == DutyReceived {
    // mark duty as being processed and start processing it
    v.Duties[v.CurrentHeight].SyncCommittee = DutyIsBeingProcessed
    go v.processDuty(v.CurrentHeight, SyncCommittee)
  }

  return
}

func (v *Validator) processDuty(height int, duty string) {
  // pretend we processed duty
  log.Println("Validator ", v.ValidatorID, ": Processed duty ", duty, " for the height ", height)

  // for updating duty status protect map
  v.DutiesLocker.Lock()

  // Use switch on duty type.
  switch {
  case duty == Proposer:
    v.Duties[height].Proposer = DutyProcessed
  case duty == Attester:
    v.Duties[height].Attester = DutyProcessed
  case duty == Aggregator:
    v.Duties[height].Aggregator = DutyProcessed
  case duty == SyncCommittee:
    v.Duties[height].SyncCommittee = DutyProcessed
  }

  // if we are the last duty that was processed in the current height increase current height
  if v.Duties[height].Proposer == DutyProcessed && v.Duties[height].Attester == DutyProcessed && v.Duties[height].Aggregator == DutyProcessed && v.Duties[height].SyncCommittee == DutyProcessed {
    log.Println("Validator ", v.ValidatorID, " moved to processing height ", v.CurrentHeight + 1)
    v.CurrentHeight++
  }

  // unlock map
  v.DutiesLocker.Unlock()

  // notify main thread to check for the next duty
  v.processingFinished <- 1

  return
}

// Checks if duty is received for processing
func (v *Validator) GetCurrentHeight() int {
  v.DutiesLocker.Lock()
  defer v.DutiesLocker.Unlock()
  return v.CurrentHeight
}

// Checks if duty is received for processing
func (v *Validator) GetDutyStatus(duty Duty) int {
  v.DutiesLocker.Lock()
  defer v.DutiesLocker.Unlock()
  // Use switch on duty type.
  switch {
  case duty.Duty == Proposer:
    return v.Duties[duty.Height].Proposer
  case duty.Duty == Attester:
    return v.Duties[duty.Height].Attester
  case duty.Duty == Aggregator:
    return v.Duties[duty.Height].Aggregator
  case duty.Duty == SyncCommittee:
    return v.Duties[duty.Height].SyncCommittee
  }

  return -1
}
