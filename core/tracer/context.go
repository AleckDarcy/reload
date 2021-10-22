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
	url     string
	server  UUID
}

type ContextMetaKey struct{}

func NewContextMeta(traceID TraceID, uuid UUID, url string) *ContextMeta {
	return &ContextMeta{traceID: traceID, uuid: uuid, url: url}
}

func NewContextMeta1(traceID TraceID, uuid UUID, url string, server UUID) *ContextMeta {
	return &ContextMeta{traceID: traceID, uuid: uuid, url: url, server: server}
}

func (c *ContextMeta) UUID() UUID {
	return c.uuid
}

func (c *ContextMeta) Url() string {
	return c.url
}

func (c *ContextMeta) ServerUUID() UUID {
	return c.server
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

	return uuid
}

func NewUUIDShort() UUID {
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		panic(err)
	}

	return uuid[len(uuid)-12:]
}
