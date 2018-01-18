// Package main provides ...
package main

import (
	"log"
	"github.com/influxdata/influxdb/client/v2"
)

const (
	username = "bonde"
	password = "bonde2017"
)

func main() {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	seed(c)
}
