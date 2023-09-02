package nexus

import (
	//"errors"
	"fmt"
	"io"
	"net/http"
)

// call is a helper function that does the actual HTTP request to the remote REST JSON API instance.
func (c *Client) call(method string, url string, payload io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("X-Auth-Token", c.Token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	fmt.Println(method + "\t" + url + "\t" + resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
