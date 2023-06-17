package validator

import (
  "fmt"
  "time"
)

// describing specific duty for specific validator
type Duty struct {
	Duty      string `json:"duty"`
	Height    int    `json:"height"`
}

// define each validator
type Validator struct {
	ValidatorID int
	Requests chan Duty
}

// validator which listens to incoming duties and processes
// the routine is ever running as for simplicity we don't implement ctx listener for
// interrupt event to shut down the validator listening process.
// The function processes request regardless height as it's unclear if the duties should be
// processed in order considering height ...
func (v *Validator) ListenForRequests(/* we could have contect here for graceful shutdown*/) {
  fmt.Println(v.ValidatorID, " started listening for incoming duties ")
  for {
    // wait for a new request
    request := <- v.Requests

    // simulating processing period ...
    time.Sleep(1 * time.Second)

    // log about the successful completion
    fmt.Println("Validator ", v.ValidatorID, " processed duty ", request.Duty, " for height ", request.Height)
  }
  return
}
