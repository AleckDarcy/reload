package tracer

import (
	"context"
	"sync/atomic"
)

type ThreadIDKey struct {}

var threadID int64

func NewThreadID() int64 {
	return atomic.AddInt64(&threadID, 1)
}

func NewContextWithThreadID(ctx context.Context) context.Context {
	return context.WithValue(ctx, ThreadIDKey{}, NewThreadID())
}
