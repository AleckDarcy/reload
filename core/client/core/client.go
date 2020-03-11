package main

import (
	"net/http"
	"net/http/cookiejar"

	"github.com/AleckDarcy/reload/core/client/data"
)

type Client struct {
	Client http.Client
}

func NewClient() *Client {
	cookieJar, _ := cookiejar.New(nil)

	return &Client{
		Client: http.Client{
			Jar: cookieJar,
		},
	}
}

func (c *Client) Call(req *data.Request) (*data.Response, error) {
	switch req.Method {
	case data.HTTPGet:
		rsp, err := c.Client.Get(req.URL)
		if err != nil {
			return nil, err
		}

		_ = rsp
	case data.HTTPPost:
		c.Client.PostForm(req.URL, req.UrlValues)
	}

	return nil, nil
}
