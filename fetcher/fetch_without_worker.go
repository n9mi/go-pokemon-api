package fetcher

import (
	"go-pokemon-api/model"
	"go-pokemon-api/utils"
)

func FetchWithoutWorker(pokemonCounts int) []model.Result {
	results := make([]model.Result, pokemonCounts)
	for i := 1; i <= pokemonCounts; i++ {
		data, err := utils.FetchData(i)
		results = append(results, model.Result{
			PokemonId: i,
			Data:      *data,
			Err:       err,
		})
	}

	return results
}
