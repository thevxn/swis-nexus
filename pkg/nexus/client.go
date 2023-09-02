package nexus

import (
	"log"
	//"go.savla.dev/swis/v5/"
)

type ClientInterface interface {
	Create(string, interface{}) bool
	Read(string, interface{}) bool
	Update(string, interface{}) bool
	Delete(string, interface{}) bool
}

type Client struct {
	BaseURL string
	Token   string
	ClientInterface
}

func (c *Client) Create(route string, payload interface{}) bool {
	_, err := call("POST", c.BaseURL + route, nil)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (c *Client) Read(route string, payload interface{}) bool {
	_, err := call("GET", c.BaseURL + route, nil)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (c *Client) Update(route string, payload interface{}) bool {
	_, err := call("PUT", c.BaseURL + route, nil)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (c *Client) Delete(route string, payload interface{}) bool {
	_, err := call("DELETE", c.BaseURL + route, nil)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

