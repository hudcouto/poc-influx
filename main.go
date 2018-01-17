// Package main provides ...
package main

import (
	"log"
	"time"
	"github.com/influxdata/influxdb/client/v2"
)

const (
	Db = "bonde"
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

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  Db,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{
		"user_first_name": "Hudson",
		"user_email": "hudson.sama@gmail.com",
		"user_id": "12",
	}

	fields := map[string]interface{}{
		"current_sign_in_ip": "192.168.10.10",
		"last_sign_in_ip": "192.168.10.10",
		"sign_in_count": "112",
	}

	pt, err := client.NewPoint("event_login", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
}
