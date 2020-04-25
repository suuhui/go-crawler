package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

//const startUrl = "http://localhost:8080/mock/www.zhenai.com/zhenghun"
const startUrl = "http://localhost:8080/mock/www.zhenai.com/zhenghun/leshan"
func main() {
	//单任务爬虫
	//e := engine.SimpleEngine{}
	//并发爬虫
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	//队列并发版爬虫
	e = engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:       startUrl,
		ParseFunc: parser.ParseCity,
	})
}
