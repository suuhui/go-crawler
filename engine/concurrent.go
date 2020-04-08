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
	Submit(request Request)
	ConfigMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.ConfigMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		doWork(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		parseResult := <-out
		for _, item := range parseResult.Items {
			log.Printf("Got item %s\n", item)
		}

		for _, r := range parseResult.Requests {
			e.Scheduler.Submit(r)
		}
		log.Println("number of goroutine: ", runtime.NumGoroutine())
	}
}

func doWork(in chan Request, out chan ParseResult) {
	go func() {
		for {
			r := <-in
			parseResult, err := worker(r)
			if err != nil {
				continue
			}

			out <- parseResult
		}
	}()
}
