package core

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/AleckDarcy/reload/runtime/html"

	"github.com/AleckDarcy/reload/core/tracer"

	"github.com/AleckDarcy/reload/core/client/data"
)

func TestHome(t *testing.T) {
	client := NewClient()

	reqs := &data.Requests{
		CookieUrl: "localhost",
		Trace:     &tracer.Trace{Id: 1},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         "http://localhost",
				MessageName: "home",
			},
		},
	}

	rsp, err := client.SendRequests(reqs)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(rsp.Body))
	t.Log(len(rsp.Trace.Records))
	t.Log(rsp.Trace)
	bytes, _ := json.Marshal(rsp.Trace)
	t.Log(string(bytes))
}

func TestHipsterShop(t *testing.T) {
	client := NewClient()

	reqs := &data.Requests{
		CookieUrl: "localhost",
		Trace:     &tracer.Trace{Id: 1},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         "http://localhost/product/OLJCESPC7Z",
				MessageName: "product",
			},
			{
				Method: data.HTTPPost,
				URL:    "http://localhost/cart",
				UrlValues: url.Values{
					"product_id": {"OLJCESPC7Z"},
					"quantity":   {"1"},
				},
				MessageName: "cart",
				Expect:      &data.ExpectedResponse{ContentType: html.ContentTypeHTML},
			},
			//{
			//	Method: data.HTTPPost,
			//	URL:    "http://localhost/cart",
			//	UrlValues: url.Values{
			//		"product_id": {"L9ECAV7KIM"},
			//		"quantity":   {"1"},
			//	},
			//	MessageName: "cart",
			//},
			//{
			//	Method:      data.HTTPGet,
			//	URL:         "http://localhost/product/L9ECAV7KIM",
			//	MessageName: "product",
			//},
			{
				Method: data.HTTPPost,
				URL:    "http://localhost/cart/checkout",
				UrlValues: url.Values{
					"email":                        {"someone@example.com"},
					"street_address":               {"1600 Amphitheatre Parkway"},
					"zip_code":                     {"94043"},
					"city":                         {"Mountain View"},
					"state":                        {"CA"},
					"country":                      {"United States"},
					"credit_card_number":           {"4432-8015-6152-0454"},
					"credit_card_expiration_month": {"1"},
					"credit_card_expiration_year":  {"2021"},
					"credit_card_cvv":              {"672"},
				},
				MessageName: "checkout",
			},
		},
	}

	rsp, err := client.SendRequests(reqs)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(rsp.Body))
	t.Log(len(rsp.Trace.Records))
	t.Log(rsp.Trace)
}
