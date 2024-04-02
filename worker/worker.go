package worker

import (
	"go-pokemon-api/model"
	"go-pokemon-api/utils"
)

type worker struct {
	id         int
	taskQueue  <-chan *WorkerInput
	resultChan chan<- *model.Result
}

func newWorker(id int, taskQueue <-chan *WorkerInput, resultChan chan<- *model.Result) worker {
	return worker{
		id:         id,
		taskQueue:  taskQueue,
		resultChan: resultChan,
	}
}

func (w *worker) start() {
	go func() {
		for n := range w.taskQueue {
			data, err := utils.FetchData(n.PokemonId)
			w.resultChan <- &model.Result{
				PokemonId: n.PokemonId,
				Data:      *data,
				Err:       err,
			}
		}
	}()
}
