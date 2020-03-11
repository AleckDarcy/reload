package core

import (
	"testing"

	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/tracer"
)

func TestHipsterShop(t *testing.T) {
	client := NewClient()

	reqs := []*data.Request{
		{
			Method:      data.HTTPGet,
			URL:         "http://localhost/product/OLJCESPC7Z",
			MessageName: "product",
			Trace:       &tracer.Trace{Id: 1},
		},
		{
			Method: data.HTTPPost,
			URL:    "http://localhost/cart",
			UrlValues: map[string][]string{
				"product_id": {"OLJCESPC7Z"},
				"quantity":   {"1"},
			},
			MessageName: "cart",
		},
		{
			Method: data.HTTPPost,
			URL:    "http://localhost/cart",
			UrlValues: map[string][]string{
				"product_id": {"L9ECAV7KIM"},
				"quantity":   {"1"},
			},
			MessageName: "cart",
		},
	}

	trace, err := client.SendRequests(reqs)
	if err != nil {
		t.Error(err)
	}

	t.Log(trace)
}
