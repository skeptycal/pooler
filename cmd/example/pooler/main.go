package main

import (
	"log"

	"github.com/skeptycal/pooler"
)

const WORKER_COUNT = 5
const JOB_COUNT = 100

func main() {
	log.Println("starting application...")
	collector := pooler.StartDispatcher(WORKER_COUNT) // start up worker pool

	for i, job := range pooler.CreateExampleJobs(JOB_COUNT) {
		collector.Work <- pooler.Work{Job: job, ID: i}
	}
}
