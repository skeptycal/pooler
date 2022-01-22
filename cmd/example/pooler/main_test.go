package main

import (
	"testing"

	"github.com/skeptycal/pooler"
)

/*  results on macOS arm test machine:

goos: darwin
goarch: arm64
pkg: github.com/skeptycal/pooler/cmd/example/pooler
BenchmarkConcurrent-8      	       1	3003465833 ns/op	    9680 B/op	     105 allocs/op
BenchmarkNonconcurrent-8   	       1	20018518667 ns/op	    2120 B/op	      49 allocs/op
*/

func BenchmarkConcurrent(b *testing.B) {
	collector := pooler.StartDispatcher(WORKER_COUNT) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, job := range pooler.CreateExampleJobs(20) {
			collector.Work <- pooler.Work{Job: job, ID: i}
		}
	}
}

func BenchmarkNonconcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, work := range pooler.CreateExampleJobs(20) {
			pooler.DoWork(work, 1)
		}
	}
}
