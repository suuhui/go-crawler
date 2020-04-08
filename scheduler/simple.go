package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(r chan engine.Request) {
	s.workerChan = r
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	//开goroutine防止concurrent engine循环等待
	go func() { s.workerChan <- request }()
}
