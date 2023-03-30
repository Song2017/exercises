package exer

import (
	"fmt"
	"sync"
)

type Task struct {
	Param  int
	Result int
}

type Pool struct {
	JobQueue    chan Task
	Results     chan Task
	WorkerCount int
	wg          sync.WaitGroup
}

func NewPool(workerCount int) *Pool {
	return &Pool{
		JobQueue:    make(chan Task),
		Results:     make(chan Task),
		WorkerCount: workerCount,
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.WorkerCount; i++ {
		go p.worker()
	}
}

func (p *Pool) worker() {
	for task := range p.JobQueue {
		task.Result = task.Param * 2
		p.Results <- task
	}
	p.wg.Done()
}

func (p *Pool) AddTask(task Task) {
	p.wg.Add(1)
	p.JobQueue <- task
}

func (p *Pool) Close() {
	close(p.JobQueue)
	p.wg.Wait()
	close(p.Results)
}

func PoolMaker() {
	pool := NewPool(3)
	pool.Start()

	// 添加任务
	pool.AddTask(Task{Param: 1})
	pool.AddTask(Task{Param: 2})
	pool.AddTask(Task{Param: 3})
	pool.AddTask(Task{Param: 4})
	pool.AddTask(Task{Param: 5})

	// 等待所有任务完成
	pool.Close()

	// 获取任务结果
	results := pool.Results
	for i := 0; i < 5; i++ {
		result := <-results
		fmt.Printf("task %d result: %d\n", result.Param, result.Result)
	}
}
