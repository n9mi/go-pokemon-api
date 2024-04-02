package test

import (
	"go-pokemon-api/fetcher"
	"testing"
)

func BenchmarkFetchWithoutWorker(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fetcher.FetchWithoutWorker(pokemonCounts)
	}
}
