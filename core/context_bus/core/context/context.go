package context

import (
	"github.com/AleckDarcy/reload/core/context_bus/code_generator"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

// RequestContext is inter-service message context
type RequestContext struct {
	lib         string // name of network API
	configureID int64
	event       *cb.EventMessage
}

func NewRequestContext(lib string, configureID int64, msg *cb.EventMessage) *RequestContext {
	return &RequestContext{
		lib:         lib,
		configureID: configureID,
		event:       msg,
	}
}

func (c *RequestContext) GetLib() string {
	return c.lib
}

func (c *RequestContext) GetConfigureID() int64 {
	return c.configureID
}

func (c *RequestContext) GetEventMessage() *cb.EventMessage {
	return c.event
}

// EventContext is the context associated with each observation
type EventContext struct {
	codebase *code_generator.CodeInfoBasic
	snapshot *cb.PrerequisiteSnapshot
}

func NewEventContext(codebase *code_generator.CodeInfoBasic, snapshot *cb.PrerequisiteSnapshot) *EventContext {
	return &EventContext{
		codebase: codebase,
		snapshot: snapshot,
	}
}

func (c *EventContext) GetCodeInfoBasic() *code_generator.CodeInfoBasic {
	return c.codebase
}

func (c *EventContext) GetPrerequisiteSnapshot() *cb.PrerequisiteSnapshot {
	return c.snapshot
}

type Context struct {
	reqCtx *RequestContext
	eveCtx *EventContext
}

func NewContext(reqCtx *RequestContext, eveCtx *EventContext) *Context {
	return &Context{
		reqCtx: reqCtx,
		eveCtx: eveCtx,
	}
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
