package validator

import "testing"

func TestAdd(t *testing.T){

    got := 1
    want := 1

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
