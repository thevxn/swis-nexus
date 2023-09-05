package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	nexus "go.savla.dev/swis-nexus/pkg/nexus"
	links "go.savla.dev/swis/v5/pkg/links"
)

var (
	// configure the client's struct parameters
	client = nexus.Client{
		BaseURL: "http://swapi.example.com",
		Token:   "xxx",
		Verbose: true,
	}

	// create a sample links.Link
	payload = links.Link{
		Name:        "test_link",
		Description: "a test link",
		URL:         "https://savla.dev",
	}
)

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	// encode the sample into a byte stream
	data, err := json.Marshal(payload)
	logError(err)

	reader := bytes.NewReader(data)
	reader2 := bytes.NewReader(data)

	// test the CRUD functions calls
	body, err := client.
		Create("/links/test", reader)
	logError(err)
	fmt.Println(string(body) + "\n")

	body, err = client.
		Read("/links/test", nil)
	logError(err)
	fmt.Println(string(body) + "\n")

	body, err = client.
		Update("/links/test", reader2)
	logError(err)
	fmt.Println(string(body) + "\n")

	body, err = client.
		Delete("/links/test", nil)
	logError(err)
	fmt.Println(string(body) + "\n")

	return

}
