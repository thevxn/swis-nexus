package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DTO-in structure for client.
type Input struct {
	// URL path or a generic route.
	Path string

	// Any JSON-tagged (required) structured data.
	Data interface{}
}

// DTO-out structure for client.
type Output struct {
	// Response status code.
	Code int `json:"-"`

	// Text API response string.
	Message string `json:"message"`

	// API timestamp.
	Timestamp int64 `json:"timestamp"`

	// Package name.
	Package string `json:"package"`

	// Data as returned by the API (mono).
	DataOne interface{} `json:"item"`

	// Data as returned by the API (poly).
	Data interface{} `json:"items"`
}

// call is a helper function that does the actual HTTP request to the remote REST JSON API instance.
func (c *Client) call(method string, input *Input, output *Output) error {
	var req *http.Request
	var err error

	// Preprocess the input data and prepare the request.
	if input.Data != nil && method != "GET" {
		// JSON-encode the input data.
		jsonData, err := json.Marshal(input.Data)
		if err != nil {
			return err
		}

		payload := bytes.NewReader(jsonData)

		// Compose a HTTP request.
		req, err = http.NewRequest(method, c.BaseURL+input.Path, payload)
		if err != nil {
			return err
		}
	} else {
		// Compose a HTTP request.
		req, err = http.NewRequest(method, c.BaseURL+input.Path, nil)
		if err != nil {
			return err
		}
	}

	// Assign the API token from client's initialization.
	req.Header.Set("X-Auth-Token", c.Token)
	req.Header.Set("Content-Type", "appliacation/json")

	// Instantialize a new HTTP client.
	client := &http.Client{}

	// Fire the request.
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Defer response body closure.
	defer resp.Body.Close()

	// Log the response details if in verbose mode.
	if c.Verbose {
		fmt.Printf("%7s \t %s \t %s\n", method, c.BaseURL+input.Path, resp.Status)
	}

	// Exit prematurely if output DTO is nil.
	if output == nil {
		return nil
	}

	// Assign the response status code.
	output.Code = resp.StatusCode

	// Read the response body.
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal the just-read response body into the output structure.
	err = json.Unmarshal(respBody, output)
	if err != nil {
		return err
	}

	return nil
}
