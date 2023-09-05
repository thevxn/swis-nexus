package nexus

import (
	"io"
)

type ClientInterface interface {
	Create(string, io.Reader) ([]byte, error)
	Read(string, io.Reader) ([]byte, error)
	Update(string, io.Reader) ([]byte, error)
	Delete(string, io.Reader) ([]byte, error)
}

type Client struct {
	BaseURL string
	Token   string
	Verbose bool
}

func (c Client) Create(route string, payload io.Reader) ([]byte, error) {
	return c.call("POST", c.BaseURL+route, payload)
}

func (c Client) Read(route string, payload io.Reader) ([]byte, error) {
	return c.call("GET", c.BaseURL+route, payload)
}

func (c Client) Update(route string, payload io.Reader) ([]byte, error) {
	return c.call("PUT", c.BaseURL+route, payload)
}

func (c Client) Delete(route string, payload io.Reader) ([]byte, error) {
	return c.call("DELETE", c.BaseURL+route, payload)
}
