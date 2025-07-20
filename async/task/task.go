package task

import "sync"

type Pool struct {
	workerNum int
	wg        sync.WaitGroup
	taskCh    chan func()
	closeCh   chan struct{}
}

func New(workerNum, bufferSize int) *Pool {
	return &Pool{
		workerNum: workerNum,
		taskCh:    make(chan func(), bufferSize),
		closeCh:   make(chan struct{}),
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
	for {
		select {
		case task := <-p.taskCh:
			task()
		case <-p.closeCh:
			return
		}
	}
}

func (p *Pool) Close() {
	close(p.closeCh)
	p.wg.Wait()
}
