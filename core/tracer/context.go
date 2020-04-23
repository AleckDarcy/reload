package tracer

import (
	"context"

	"github.com/hashicorp/go-uuid"
)

var ServiceUUID = NewUUID()

type TraceID = int64
type UUID = string

type ContextMeta struct {
	traceID TraceID
	uuid    UUID
}

type ContextMetaKey struct{}

func NewContextMeta(traceID TraceID, uuid UUID) *ContextMeta {
	return &ContextMeta{traceID: traceID, uuid: uuid}
}

func (c *ContextMeta) UUID() UUID {
	return c.uuid
}

func NewContextWithContextMeta(ctx context.Context, c *ContextMeta) context.Context {
	return context.WithValue(ctx, ContextMetaKey{}, c)
}

func NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextMetaKey{}, &ContextMeta{})
}

func NewUUID() UUID {
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		panic(err)
	}

	return UUID(uuid)
}
