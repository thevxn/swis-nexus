package nexus

import (
	"fmt"
)

type ClientInterface interface {
	Get(*Input, *Output) error
	Post(*Input, *Output) error
	Patch(*Input, *Output) error
	Put(*Input, *Output) error
	Delete(*Input, *Output) error
}

// nexus client structure
type Client struct {
	// Base API's URL, hostname.
	BaseURL string

	// API authentication token.
	Token string

	// Verbose mode to print the call details.
	Verbose bool
}

// Compose and return a new Client instance.
func NewClient(baseURL, token string) *Client {
	if baseURL == "" {
		fmt.Println("baseURL string has to be specified")
		return nil
	}

	// Return the pointer to Client struct.
	return &Client{
		BaseURL: baseURL,
		Token:   token,
		Verbose: false,
	}
}

// GET HTTP method.
func (c *Client) Get(input *Input, output *Output) error {
	return c.call("GET", input, output)
}

// POST HTTP method metacall.
func (c *Client) Post(input *Input, output *Output) error {
	return c.call("POST", input, output)
}

// PUT HTTP method metacall.
func (c *Client) Put(input *Input, output *Output) error {
	return c.call("PUT", input, output)
}

// PATCH HTTP method metacall.
func (c *Client) Patch(input *Input, output *Output) error {
	return c.call("PATCH", input, output)
}

// DELETE HTTP method metacall.
func (c *Client) Delete(input *Input, output *Output) error {
	return c.call("DELETE", input, output)
}
