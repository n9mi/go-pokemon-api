package fetcher

import (
	"go-pokemon-api/model"
	"go-pokemon-api/worker"
	"runtime"
)

func FetchWithWorker(pokemonCounts int) []model.Result {
	// Generate task
	tasks := make([]worker.WorkerInput, pokemonCounts)
	for i := 1; i <= pokemonCounts; i++ {
		tasks[i-1] = worker.WorkerInput{
			PokemonId: i,
		}
	}

	// Create and start worker
	numWorker := runtime.GOMAXPROCS(0)
	workerPool := worker.NewWorkerPool(numWorker)
	workerPool.Start()

	// Assign tasks to workers
	for _, task := range tasks {
		workerPool.Submit(&task)
	}

	// Collect results and handle error
	var results []model.Result
	for i := 0; i < pokemonCounts; i++ {
		result := workerPool.GetResult()
		results = append(results, *result)
	}

	return results
}
