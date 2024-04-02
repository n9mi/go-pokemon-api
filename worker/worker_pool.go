package worker

import "go-pokemon-api/model"

type WorkerPool struct {
	TaskQueue   chan *WorkerInput
	ResultChan  chan *model.Result
	WorkerCount int
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		TaskQueue:   make(chan *WorkerInput, workerCount),
		ResultChan:  make(chan *model.Result, workerCount),
		WorkerCount: workerCount,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.WorkerCount; i++ {
		worker := newWorker(i, wp.TaskQueue, wp.ResultChan)
		worker.start()
	}
}

func (wp *WorkerPool) Submit(workerInput *WorkerInput) {
	wp.TaskQueue <- workerInput
}

func (wp *WorkerPool) GetResult() *model.Result {
	return <-wp.ResultChan
}
