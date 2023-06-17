package validator

// describing specific duty for specific validator
type Duty struct {
	Duty      string `json:"duty"`
	Height    int    `json:"height"`
}

// define each validator
type Validator struct {
	ValidatorID int
	requests chan Duty
}

// validator which listens to incoming duties and processes
// the routine is ever running as for simplicity we don't implement ctx listener for
// interrupt event to shut down the validator listening process
func (v *Validator) ListenForRequests(wg *sync.WaitGroup) {
  for {
    // wait for a new request
    request <- v.requests

    // simulating processing period ...
    time.Sleep(1 * time.Duration)

    // log about the successful completion
    fmt.Println("processed request ", request.Duty, " for height ", request.Height)
  }
  return
}
