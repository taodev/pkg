package task

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPool(t *testing.T) {
	workerNum := 3
	bufferSize := 10
	pool := New(workerNum, bufferSize)

	assert.Equal(t, workerNum, pool.workerNum, "workerNum should be set correctly")
	assert.Equal(t, bufferSize, cap(pool.taskCh), "taskCh capacity should be set correctly")
}

func TestPool_StartAndDo(t *testing.T) {
	pool := New(2, 5)

	var counter int32
	 taskCount := 10

	// 启动工作池
	pool.Start()

	// 提交任务
	for i := 0; i < taskCount; i++ {
		pool.Do(func() {
			atomic.AddInt32(&counter, 1)
			time.Sleep(10 * time.Millisecond) // 模拟任务执行时间
		})
	}

	// 给足够时间让所有任务完成
	pool.Close()

	// 验证所有任务都已执行
	assert.Equal(t, int32(taskCount), atomic.LoadInt32(&counter), "all tasks should be executed")
}

func TestPool_Close(t *testing.T) {
	pool := New(2, 5)
	pool.Start()

	var counter int32
	for i := 0; i < 3; i++ {
		pool.Do(func() {
			atomic.AddInt32(&counter, 1)
		})
	}

	pool.Close()

	// 验证关闭后无法提交任务
	defer func() {
		r := recover()
		assert.NotNil(t, r, "expected panic when submitting to closed pool")
	}()
	pool.Do(func() {})
}

func TestPool_ParallelTasks(t *testing.T) {
	pool := New(4, 20)
	defer pool.Close()
	pool.Start()

	var counter int32
	 taskCount := 100
	wg := &sync.WaitGroup{}
	wg.Add(taskCount)

	for i := 0; i < taskCount; i++ {
		pool.Do(func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		})
	}

	wg.Wait()

	assert.Equal(t, int32(taskCount), atomic.LoadInt32(&counter), "all parallel tasks should be executed")
}