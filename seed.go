package main

import (
	"log"
	"time"
	"github.com/influxdata/influxdb/client/v2"
	"math/rand"
	"fmt"
)

func seed(clnt client.Client) {
	sampleSize := 1000

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "bonde",
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sampleSize; i++ {
		regions := []string{"South America", "Central America", "Northern America", "Australia", "Western Africa"}
		tags := map[string]string{
			"ip":    fmt.Sprintf("192.168.172.%d", rand.Intn(1000)),
			"host":   fmt.Sprintf("host%d", rand.Intn(1000)),
			"region": regions[rand.Intn(len(regions))],
		}

		os := []string{"OSX", "Windows", "Linux"}
		browsers := []string{"Google Chrome", "Firefox", "Safari", "Opera"}
		fields := map[string]interface{}{
			"browser": browsers[rand.Intn(len(browsers))],
			"os": os[rand.Intn(len(os))],
		}

		timein := time.Now().Local().Add(time.Hour +
		time.Minute * time.Duration(i+1) +
		time.Second * time.Duration(i))

		pt, err := client.NewPoint("login", tags, fields, timein)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	if err := clnt.Write(bp); err != nil {
		log.Fatal(err)
	}
}
