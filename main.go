package main

import (
	"swis-nexus/pkg/nexus"
)

func main() {
	client := nexus.Client{
		BaseURL: "swapi.savla.su",
		Token: "xxx",
	}
	
	client.Create("/users", nil)
	client.Read("/users", nil)
	client.Update("/users", nil)
	client.Delete("/users", nil)
}
