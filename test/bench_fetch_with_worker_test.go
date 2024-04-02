package test

import (
	"go-pokemon-api/fetcher"
	"testing"
)

const pokemonCounts = 10

func BenchmarkFetchWithWorker(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fetcher.FetchWithWorker(pokemonCounts)
	}
}
