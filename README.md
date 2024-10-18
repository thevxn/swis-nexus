# swis-nexus

[![Go Reference](https://pkg.go.dev/badge/go.savla.dev/swis-nexus.svg)](https://pkg.go.dev/go.savla.dev/swis-nexus)
[![Go Report Card](https://goreportcard.com/badge/go.savla.dev/swis-nexus)](https://goreportcard.com/report/go.savla.dev/swis-nexus)

A simple client/connector for [swis-api](https://github.com/savla-dev/swis-api) RESTful JSON API. 

A complete example implementation can be find in [cmd/swis-nexus/main.go](/cmd/swis-nexus/main.go) file.

### import a usage

```shell
go get go.vxn.dev/swis-nexus/pkg/nexus
```

```go
package main

import (
    "fmt"

    "go.vxn.dev/swis-nexus/pkg/nexus"
)

var (
    client *nexus.Client

    baseURL = "https://swapi.example.com"
    token   = "xxx"
    verbose = true
)

func main() {
    // Fetch a new nexus.Client instance.
    client = nexus.NewClient(baseURL, token)

    // Enable the verbose mode.
    client.Verbose = true

    // Compose a DTO-in object.
    input := &nexus.Input{
        Path: "/users",
        Data: nil,
    }

    // Prepare a DTO-out object.
    output := &nexus.Output{}

    // Execute the API call.
    if err := client.Get(input, output); err != nil {
        fmt.Println(err)
    }

    // Print the response code and response message.
    fmt.Printf("%4d: %s", output.Code, output.Message)
}
```
