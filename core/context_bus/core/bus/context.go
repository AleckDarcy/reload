package bus

import (
	"github.com/AleckDarcy/reload/core/context_bus/code_generator"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

// RequestContext is inter-service message context
type RequestContext struct {
	lib         string // name of network API
	configureID int64
	attrs       *cb.Attributes
}

func NewRequestContext(lib string, configureID int64, attrs *cb.Attributes) *RequestContext {
	return &RequestContext{
		lib:         lib,
		configureID: configureID,
		attrs:       attrs,
	}
}

func (c *RequestContext) GetLib() string {
	return c.lib
}

func (c *RequestContext) GetConfigureID() int64 {
	return c.configureID
}

func (c *RequestContext) GetAttrs() *cb.Attributes {
	return c.attrs
}

// EventContext is the context associated with each observation
type EventContext struct {
	codebase *code_generator.CodeInfoBasic
}

type Context struct {
	reqCtx *RequestContext
	eveCtx *EventContext
}

func (c *Context) GetRequestContext() *RequestContext {
	return c.reqCtx
}

// SetRequestContext is written by network APIs on receiving requests. e.g., rest, rpc
// contains the request-wise static values. e.g., session-id, token
func (c *Context) SetRequestContext(reqCtx *RequestContext) *Context {
	newC := *c
	newC.reqCtx = reqCtx

	return &newC
}

func (c *Context) GetEventContext() *EventContext {
	return c.eveCtx
}

// SetEventContext is written by generated code when user or library submit their observations
// contains the static code base information
func (c *Context) SetEventContext(eveCtx *EventContext) *Context {
	newC := *c
	newC.eveCtx = eveCtx

	return &newC
}
