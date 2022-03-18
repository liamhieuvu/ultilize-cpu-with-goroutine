package main

import (
	"log"
	"sync"
	"time"
)

type item struct {
	ID       string
	Duration time.Duration
}

func main() {
	items := []item{{"A", 1}, {"B", 2}, {"C", 1}, {"D", 3}, {"E", 2}}
	n := 2
	start := time.Now()
	for s := 0; s < len(items); s += n {
		e := s + 2
		if s+n > len(items) {
			e = len(items)
		}
		processN(items[s:e])
	}
	log.Print("processed time: ", time.Since(start))
}

func processN(items []item) {
	var wg sync.WaitGroup
	wg.Add(len(items))
	for _, data := range items {
		go func(i item) {
			defer wg.Done()
			process(i)
		}(data)
	}
	wg.Wait()
}

func process(i item) {
	log.Print(i.ID, " start")
	time.Sleep(i.Duration * time.Second)
	log.Print(i.ID, " end")
}
