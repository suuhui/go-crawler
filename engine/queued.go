package engine

import (
	"log"
)

type QueuedEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *QueuedEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i ++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d %v", itemCount, item)
			itemCount++
		}

		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(out chan ParseResult, scheduler Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			scheduler.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
