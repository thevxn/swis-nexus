# swis-nexus

[![Go Reference](https://pkg.go.dev/badge/go.savla.dev/swis-nexus.svg)](https://pkg.go.dev/go.savla.dev/swis-nexus)
[![Go Report Card](https://goreportcard.com/badge/go.savla.dev/swis-nexus)](https://goreportcard.com/report/go.savla.dev/swis-nexus)

A simple client/connector for [swis-api](https://github.com/savla-dev/swis-api) RESTful JSON API. 

An example implementation can be seen in [cmd/main.go](/cmd/main.go) file.

### import

```shell
go install go.savla.dev/swis-nexus@latest
```

```go
package main

import (
    "go.savla.dev/swis-nexus/pkg/nexus"
)

func main() {
    client := nexus.Client{
        BaseURL: "https://swapi.example.com",
        Token:   "xxx",
    }

    client.Read("/users", nil)
}
```
