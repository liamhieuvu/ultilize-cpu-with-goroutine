package main

import (
	"log"
	"time"
)

type item struct {
	ID       string
	Duration time.Duration
}

func main() {
	items := []item{{"A", 1}, {"B", 2}, {"C", 1}, {"D", 3}, {"E", 2}}
	n := 2
	data := make(chan item, len(items))
	results := make(chan error, len(items))
	start := time.Now()

	for w := 1; w <= n; w++ {
		go worker(data, results)
	}

	for _, d := range items {
		data <- d
	}
	close(data)

	for i := 1; i <= len(items); i++ {
		<-results
	}

	log.Print("processed time: ", time.Since(start))
}

func worker(data <-chan item, results chan<- error) {
	for d := range data {
		process(d)
		results <- nil
	}
}

func process(i item) {
	log.Print(i.ID, " start")
	time.Sleep(i.Duration * time.Second)
	log.Print(i.ID, " end")
}
