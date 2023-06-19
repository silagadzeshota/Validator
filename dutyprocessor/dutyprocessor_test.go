package dutyprocessor

import (
  "time"
  "testing"
  "dutyprocessor/validator"
)

func TestValidatorExists(t *testing.T){
    // create duty processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "21", Duty: validator.Proposer, Height: 0})

    // should not exist
    if processor.ValidatorExists("1") == true {
      t.Errorf("validator with id 1 should not exist")
    }

    // should not exist
    if processor.ValidatorExists("21") == false {
      t.Errorf("validator with id 21 should  exist")
    }
}

func TestIncorrectValidatorID(t *testing.T){
    // create duty processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "21s", Duty: validator.Proposer, Height: 0})

    // should not exist
    if processor.ValidatorExists("21s") == true {
      t.Errorf("validator with id 21s should not exist")
    }
}

func TestIncorrectHeight(t *testing.T){
    // create duty processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: -2})

    // should not exist
    if processor.ValidatorExists("1") == true {
      t.Errorf("validator with id 1 should exist")
    }
}

func TestIncorrectDuty(t *testing.T){
    // create duty processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: "someotherduty", Height: 2})

    // should not exist
    if processor.ValidatorExists("1") == true {
      t.Errorf("validator with id 1 should exist")
    }
}

func TestValidatorProcessedDuty(t *testing.T){
    // as we have goroutines we need to use sleep for simplicity so not to wait the sum of all sleeps we run tests in parallel :)
    t.Parallel()

    // create main processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: 0})

    // should not exist
    time.Sleep(1 * time.Second)

    if processor.ValidatorExists("1") == false {
      t.Errorf("validator with id 1 should exist")
    }

    status := processor.GetValidator("1").GetDutyStatus(validator.Duty{Duty: validator.Proposer, Height: 0})
    if status != 3 {
        t.Errorf("validator should have processed duty by now %d",status)
    }
}

func TestValidatorProcessesDuties(t *testing.T){
    // as we have goroutines we need to use sleep for simplicity so not to wait the sum of all sleeps we run tests in parallel :)
    t.Parallel()

    // create main processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Aggregator, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.SyncCommittee, Height: 0})

    // should not exist
    time.Sleep(1 * time.Second)

    if processor.ValidatorExists("1") == false {
      t.Errorf("validator with id 1 should exist")
    }

    status1 := processor.GetValidator("1").GetDutyStatus(validator.Duty{Duty: validator.Proposer, Height: 0})
    status2 := processor.GetValidator("1").GetDutyStatus(validator.Duty{Duty: validator.Attester, Height: 0})
    status3 := processor.GetValidator("1").GetDutyStatus(validator.Duty{Duty: validator.Aggregator, Height: 0})
    status4 := processor.GetValidator("1").GetDutyStatus(validator.Duty{Duty: validator.SyncCommittee, Height: 0})
    if status1 + status2 + status3 + status4 != 12 {
        t.Errorf("validator should have processed all the duties by now %d", status1 + status2 + status3 + status4)
    }
}

func TestValidatorDoesntProcessUnrequestedDuty(t *testing.T){
    // as we have goroutines we need to use sleep for simplicity so not to wait the sum of all sleeps we run tests in parallel :)
    t.Parallel()

    // create main processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Aggregator, Height: 0})

    // should not exist
    time.Sleep(1 * time.Second)

    if processor.ValidatorExists("1") == false {
      t.Errorf("validator with id 1 should exist")
    }

    status1 := processor.GetValidator("1").GetDutyStatus(validator.Duty{Duty: validator.Proposer, Height: 0})
    status2 := processor.GetValidator("1").GetDutyStatus(validator.Duty{Duty: validator.Attester, Height: 0})
    status3 := processor.GetValidator("1").GetDutyStatus(validator.Duty{Duty: validator.Aggregator, Height: 0})
    status4 := processor.GetValidator("1").GetDutyStatus(validator.Duty{Duty: validator.SyncCommittee, Height: 0})
    if status1 + status2 + status3 + status4 != 9 {
        t.Errorf("validator didn't process duties correctly %d", status1 + status2 + status3 + status4)
    }
}

func TestValidatorDoesnIncreaseHeightUntillAllDutiesProcessed(t *testing.T){
    // as we have goroutines we need to use sleep for simplicity so not to wait the sum of all sleeps we run tests in parallel :)
    t.Parallel()

    // create main processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Aggregator, Height: 0})

    // should not exist
    time.Sleep(1 * time.Second)

    if processor.ValidatorExists("1") == false {
      t.Errorf("validator with id 1 should exist")
    }

    if processor.GetValidator("1").GetCurrentHeight() != 0 {
        t.Errorf("Increased height without processing all the duties")
    }
}

func TestValidatorIncreaseHeightAfterProcessingAllDuties(t *testing.T){
    // as we have goroutines we need to use sleep for simplicity so not to wait the sum of all sleeps we run tests in parallel :)
    t.Parallel()

    // create main processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Aggregator, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.SyncCommittee, Height: 0})

    // should not exist
    time.Sleep(1 * time.Second)

    if processor.GetValidator("1").GetCurrentHeight() != 1 {
        t.Errorf("Validator didn't increase height")
    }
}

func TestValidatorDoesntProcessHigherHeightDuties(t *testing.T){
    // as we have goroutines we need to use sleep for simplicity so not to wait the sum of all sleeps we run tests in parallel :)
    t.Parallel()

    // create main processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: 1})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Attester, Height: 1})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Aggregator, Height: 1})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.SyncCommittee, Height: 1})

    // should not exist
    time.Sleep(300 * time.Millisecond)

    if processor.GetValidator("1").GetCurrentHeight() != 0 {
        t.Errorf("Validator increased height without processing all duties")
    }

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Aggregator, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.SyncCommittee, Height: 0})

    // should not exist
    time.Sleep(300 * time.Millisecond)
    if processor.GetValidator("1").GetCurrentHeight() != 2 {
        t.Errorf("Validator didn't increased height after processing all duties")
    }
}

func TestValidatorDoesntProcessHigherHeightDuty(t *testing.T){
    // as we have goroutines we need to use sleep for simplicity so not to wait the sum of all sleeps we run tests in parallel :)
    t.Parallel()

    // create main processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Aggregator, Height: 0})

    // should not exist
    time.Sleep(300 * time.Millisecond)

    if processor.GetValidator("1").GetCurrentHeight() != 0 {
        t.Errorf("Validator increased height without processing all duties")
    }

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: 1})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Attester, Height: 1})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Aggregator, Height: 1})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.SyncCommittee, Height: 1})

    // should not exist
    time.Sleep(300 * time.Millisecond)
    if processor.GetValidator("1").GetCurrentHeight() != 0 {
        t.Errorf("Validator didn't increased height after processing all duties")
    }
}

func TestMultipleValidators(t *testing.T){
    // as we have goroutines we need to use sleep for simplicity so not to wait the sum of all sleeps we run tests in parallel :)
    t.Parallel()

    // create main processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Aggregator, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.SyncCommittee, Height: 0})

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Aggregator, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.SyncCommittee, Height: 0})

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "3", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "3", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "3", Duty: validator.Aggregator, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "3", Duty: validator.SyncCommittee, Height: 1})

    // should not exist
    time.Sleep(300 * time.Millisecond)

    if processor.GetValidator("1").GetCurrentHeight() != 1 {
        t.Errorf("Validator didn't increase height")
    }
    if processor.GetValidator("2").GetCurrentHeight() != 1 {
        t.Errorf("Validator didn't increase height")
    }
    if processor.GetValidator("3").GetCurrentHeight() != 0 {
        t.Errorf("Validator didn't increase height")
    }
}


func TestMultipleValidatorsNotIncreasingHeight(t *testing.T){
    // as we have goroutines we need to use sleep for simplicity so not to wait the sum of all sleeps we run tests in parallel :)
    t.Parallel()

    // create main processor
    processor := NewDutyProcessor()

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.Aggregator, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "1", Duty: validator.SyncCommittee, Height: 221})

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Proposer, Height: 2})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Attester, Height: 2})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Aggregator, Height: 2})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.SyncCommittee, Height: 2})

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Proposer, Height: 1})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Attester, Height: 1})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Aggregator, Height: 1})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.SyncCommittee, Height: 1})

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.Aggregator, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "2", Duty: validator.SyncCommittee, Height: 0})

    // create validator with id 21
    processor.ProcessRequest(DutyRequest{Validator: "3", Duty: validator.Proposer, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "3", Duty: validator.Attester, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "3", Duty: validator.Aggregator, Height: 0})
    processor.ProcessRequest(DutyRequest{Validator: "3", Duty: validator.SyncCommittee, Height: 1})

    // should not exist
    time.Sleep(300 * time.Millisecond)

    if processor.GetValidator("1").GetCurrentHeight() != 0 {
        t.Errorf("Validator didn't increase height")
    }
    if processor.GetValidator("2").GetCurrentHeight() != 3 {
        t.Errorf("Validator didn't increase height")
    }
    if processor.GetValidator("3").GetCurrentHeight() != 0 {
        t.Errorf("Validator didn't increase height")
    }
}
