package nexus

import (
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
	Token string
	ClientInterface
}

func (c *Client) Create(route string, payload interface{}) bool {
	return call("GET", route, nil)
}

func (c *Client) Read(route string, payload interface{}) bool {
	return false
}

func (c *Client) Update(route string, payload interface{}) bool {
	return false
}

func (c *Client) Delete(route string, payload interface{}) bool {
	return false
}

func call(method string, route string, payload interface{}) bool {
	return false
}


