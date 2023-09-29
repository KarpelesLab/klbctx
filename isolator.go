package klbctx

import (
	"context"
	"time"
)

type isolator struct {
	context.Context
}

// NewIsolator returns a context that is a child of the passed context,
// but will not be cancelled by any condition such as deadlines or cancellations.
func NewIsolator(parent context.Context) context.Context {
	return &isolator{parent}
}

func (i *isolator) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (i *isolator) Done() <-chan struct{} {
	return nil
}

func (i *isolator) Err() error {
	return nil
}
