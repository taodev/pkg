package task

import "golang.org/x/sync/errgroup"

type Pool struct {
	workerNum int
	wg        errgroup.Group
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
		p.wg.Go(p.worker)
	}
}

func (p *Pool) worker() error {
	var taskFn func()
	var ok bool
	for {
		if taskFn, ok = <-p.taskCh; !ok {
			return nil
		}
		taskFn()
	}
}

func (p *Pool) Close() error {
	close(p.taskCh)
	return p.wg.Wait()
}

func (p *Pool) Do(task func()) {
	p.taskCh <- task
}
