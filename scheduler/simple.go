package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) GetWorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	//开goroutine防止concurrent engine循环等待
	//每个request都新开一个goroutine，因此goroutine数量很大
	go func() { s.workerChan <- request }()
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
