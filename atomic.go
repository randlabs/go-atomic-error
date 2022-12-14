package exterror

import (
	"sync"
	"time"
)

// -----------------------------------------------------------------------------

// AtomicError is a thread-safe error object. It also implements Context behavior
type AtomicError struct {
	mtx      sync.RWMutex
	err      error
	done     chan struct{}
	doneOnce sync.Once
}

// -----------------------------------------------------------------------------

// NewAtomicError creates a new thread safe error object.
func NewAtomicError() *AtomicError {
	e := &AtomicError{
		mtx:      sync.RWMutex{},
		done:     make(chan struct{}),
		doneOnce: sync.Once{},
	}
	return e
}

// Set stores the passed error if the current is nil and completes the context.
func (x *AtomicError) Set(err error) bool {
	changed := false

	if err != nil {
		x.mtx.Lock()
		if x.err == nil {
			changed = true
			x.err = err
		}
		x.mtx.Unlock()

		if changed {
			x.doneOnce.Do(func() {
				close(x.done)
			})
		}
	}

	return changed
}

// Deadline returns the time when work done on behalf of this context
// should be canceled. There is no Deadline for a AtomicError.
func (*AtomicError) Deadline() (deadline time.Time, ok bool) {
	return
}

// Done returns a channel that's closed when work done on behalf of this
// context should be canceled.
func (x *AtomicError) Done() <-chan struct{} {
	return x.done
}

// Value returns the value associated with this context for key, or nil
// if no value is associated with key.
func (*AtomicError) Value(_ interface{}) interface{} {
	return nil
}

// Err returns nil if Done is not yet closed.
// If Done is closed, Err returns a non-nil error explaining why:
// Canceled if the context was canceled
// or DeadlineExceeded if the context's deadline passed.
// After Err returns a non-nil error, successive calls to Err return the same error.
func (x *AtomicError) Err() error {
	x.mtx.RLock()
	defer x.mtx.RUnlock()

	return x.err
}
