package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

const startUrl = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main() {
	//e := engine.SimpleEngine{}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 1,
	}
	e.Run(engine.Request{
		Url:       startUrl,
		ParseFunc: parser.ParseCityList,
	})
}
