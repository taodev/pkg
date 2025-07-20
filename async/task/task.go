package task

import "sync"

type Pool struct {
	workerNum int
	wg        sync.WaitGroup
	taskCh    chan func()
}

func New(workerNum, bufferSize int) *Pool {
	return &Pool{
		workerNum: workerNum,
		taskCh:    make(chan func(), bufferSize),
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.workerNum; i++ {
		p.wg.Add(1)
		go p.worker()
	}
}

func (p *Pool) worker() {
	defer p.wg.Done()
	var taskFn func()
	var ok bool
	for {
		if taskFn, ok = <-p.taskCh; !ok {
			return
		}
		taskFn()
	}
}

func (p *Pool) Close() {
	close(p.taskCh)
	p.wg.Wait()
}

func (p *Pool) Do(task func()) {
	p.taskCh <- task
}
