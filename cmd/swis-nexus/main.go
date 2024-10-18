package main

import (
	"fmt"

	nexus "go.vxn.dev/swis-nexus/pkg/nexus"
	links "go.vxn.dev/swis/v5/pkg/links"
)

var (
	// Declare the client's struct.
	client *nexus.Client

	// Define the runtime properties.
	baseURL = "https://swapi.example.com"
	token   = "xxx"
	verbose = true

	// Create a sample links.Link data payload.
	payload = links.Link{
		ID:          "test_link",
		Name:        "test_link",
		Description: "a test link",
		URL:         "https://vxn.dev",
		Active:      true,
	}
)

// A wrapper function to catch the possible error.
func wrapError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func main() {
	// Fetch new client instance.
	client = nexus.NewClient(baseURL, token)

	// Enable the verbose mode explicitly.
	client.Verbose = true

	// Prepare the API call input.
	input := &nexus.Input{
		Path: "/links",
		Data: payload,
	}

	// Prepare the API call output.
	output := &nexus.Output{Data: &links.Link{}}

	//
	// Perform the example call sequence (create and fetch).
	//

	// Create a new link.
	if !wrapError(client.Post(input, nil)) {
		return
	}

	input.Path = "/links/test_link"

	// Fetch such link's data.
	if !wrapError(client.Get(input, output)) {
		return
	}

	// Delete the link.
	if !wrapError(client.Delete(input, nil)) {
		return
	}

	// Print fetched data from API.
	link, ok := output.Data.(*links.Link)
	if !ok {
		fmt.Println("could not assert the data type to received data")
		return
	}

	if output.Code == 200 {
		fmt.Printf("fetched link's URL: %s\n", link.URL)
	}
	return
}
