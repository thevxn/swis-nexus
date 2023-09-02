package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"go.savla.dev/swis/v5/pkg/links"
	//"go.savla.dev/swis-nexus/pkg/nexus"

	"swis-nexus/pkg/nexus"
)

func main() {
	// configure the client's struct parameters
	client := nexus.Client{
		BaseURL: "http://swapi.example.com",
		Token:   "xxx",
		Verbose: true,
	}

	// create a sample links.Link
	payload := links.Link{
		Name:        "test_link",
		Description: "a test link",
		URL:         "https://savla.dev",
	}

	// encode JSON into a byte stream
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	reader2 := bytes.NewReader(data)

	// test the CRUD fucntions calls
	body, ok := client.Create("/links/test", reader)
	if ok {
		fmt.Println(string(body) + "\n")
	}

	body, ok = client.Read("/links/test", nil)
	if ok {
		fmt.Println(string(body) + "\n")
	}

	body, ok = client.Update("/links/test", reader2)
	if ok {
		fmt.Println(string(body) + "\n")
	}

	body, ok = client.Delete("/links/test", nil)
	if ok {
		fmt.Println(string(body) + "\n")
	}

}
