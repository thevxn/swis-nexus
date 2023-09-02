package nexus

import (
	//"errors"
	"io"
	"log"
	"net/http"
)

// getResponseBody does the actual HTTP GET request and returns read response body contents.
func call(method string, url string, payload io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return []byte{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	log.Println("status: " + resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
