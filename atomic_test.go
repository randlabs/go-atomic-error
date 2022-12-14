package exterror_test

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/randlabs/go-exterror"
)

// -----------------------------------------------------------------------------

func TestAtomicError(t *testing.T) {
	wg := sync.WaitGroup{}

	err := exterror.NewAtomicError()

	// Simulate two go-routines setting an error simultaneously
	wg.Add(2)
	go func() {
		err.Set(errors.New("error 1"))
		wg.Done()
	}()

	go func() {
		err.Set(errors.New("error 2"))
		wg.Done()
	}()

	wg.Wait()

	if err.Err().Error() != "error 1" && err.Err().Error() != "error 2" {
		t.FailNow()
	}
}

func TestAtomicErrorContext(t *testing.T) {
	err := exterror.NewAtomicError()

	go func() {
		err.Set(errors.New("error"))
	}()

	select {
	case <-err.Done():
	case <-time.After(5 * time.Second):
		t.FailNow()
	}
}
