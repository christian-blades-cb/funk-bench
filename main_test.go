package main

import (
	"github.com/thoas/go-funk"
	"math/rand"
	"testing"
)

const (
	seed      = 918234565
	sliceSize = 3614562
)

func sliceGenerator(size uint, r *rand.Rand) (out []int64) {
	for i := uint(0); i < size; i++ {
		out = append(out, rand.Int63())
	}
	return
}

func BenchmarkContains(b *testing.B) {
	r := rand.New(rand.NewSource(seed))
	testData := sliceGenerator(sliceSize, r)
	what := r.Int63()

	b.Run("Idiomatic", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			idiomaticContains(what, testData)
		}
	})

	b.Run("funk", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			funk.Contains(testData, what)
		}
	})
}

func idiomaticContains(what int64, slice []int64) bool {
	for _, item := range slice {
		if what == item {
			return true
		}
	}
	return false
}
