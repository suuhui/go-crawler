package engine

import (
	"log"
	"runtime"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

//如果将Scheduler定义在scheduler中，会导致循环import
type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	GetWorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		doWork(e.Scheduler.GetWorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		parseResult := <-out
		for _, item := range parseResult.Items {
			log.Printf("Got item #%d %v\n", itemCount, item)
			itemCount++
		}

		for _, r := range parseResult.Requests {
			e.Scheduler.Submit(r)
		}
		log.Println("number of goroutine: ", runtime.NumGoroutine())
	}
}

func doWork(in chan Request, out chan ParseResult, notifier ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			r := <-in
			parseResult, err := worker(r)
			if err != nil {
				continue
			}

			out <- parseResult
		}
	}()
}
