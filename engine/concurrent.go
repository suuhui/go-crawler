package engine

import (
	"crawler/model"
	"crypto/md5"
	"encoding/hex"
	"log"
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

var foundUrls = make(map[string]bool)
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		doWork(e.Scheduler.GetWorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		hasFound(r.Url)
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		parseResult := <-out
		for _, item := range parseResult.Items {
			if _, ok := item.(model.UserProfile); ok {
				log.Printf("Got profile #%d %v\n", itemCount, item)
				itemCount++
			}
		}

		for _, r := range parseResult.Requests {
			if hasFound(r.Url) {
				continue
			}
			e.Scheduler.Submit(r)
		}
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

func hasFound(url string) bool {
	encodeStr := md5Encode(url)
	if _, ok := foundUrls[encodeStr]; ok {
		return true
	}
	foundUrls[encodeStr] = true
	return false
}

func md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
