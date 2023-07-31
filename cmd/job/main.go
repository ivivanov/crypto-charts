package main

import (
	"log"

	"github.com/ivivanov/crypto-charts/pkg/job"
)

const ENV_PATH = "./.env"

func main() {
	job, err := job.NewJob(ENV_PATH)
	if err != nil {
		log.Fatal(err)
	}

	err = job.Run()
	if err != nil {
		log.Fatal(err)
	}
}
