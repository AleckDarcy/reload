package observation

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
	
	"testing"
)

func TestA(t *testing.T) {
	t.Log(cb.Payload{})
}