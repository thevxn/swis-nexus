package nexus

import (
	"io"
	"log"
)

type ClientInterface interface {
	Create(string, io.Reader) ([]byte, bool)
	Read(string, io.Reader) ([]byte, bool)
	Update(string, io.Reader) ([]byte, bool)
	Delete(string, io.Reader) ([]byte, bool)
}

type Client struct {
	BaseURL string
	Token   string
	Verbose bool
}

func (c *Client) Create(route string, payload io.Reader) ([]byte, bool) {
	body, err := c.call("POST", c.BaseURL+route, payload)
	if err != nil {
		log.Fatal(err)
		return []byte{}, false
	}

	return body, true
}

func (c *Client) Read(route string, payload io.Reader) ([]byte, bool) {
	body, err := c.call("GET", c.BaseURL+route, payload)
	if err != nil {
		log.Fatal(err)
		return []byte{}, false
	}

	return body, true
}

func (c *Client) Update(route string, payload io.Reader) ([]byte, bool) {
	body, err := c.call("PUT", c.BaseURL+route, payload)
	if err != nil {
		log.Fatal(err)
		return []byte{}, false
	}

	return body, true
}

func (c *Client) Delete(route string, payload io.Reader) ([]byte, bool) {
	body, err := c.call("DELETE", c.BaseURL+route, payload)
	if err != nil {
		log.Fatal(err)
		return []byte{}, false
	}

	return body, true
}
