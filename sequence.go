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

	start := time.Now()
	for _, data := range items {
		process(data)
	}
	log.Print("processed time: ", time.Since(start))
}

func process(i item) {
	log.Print(i.ID, " start")
	time.Sleep(i.Duration * time.Second)
	log.Print(i.ID, " end")
}
