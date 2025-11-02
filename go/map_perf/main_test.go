package main

import (
	"crypto/rand"
	"github.com/stretchr/testify/require"
	mathRand "math/rand/v2"
	"slices"
	"testing"
)

type RunInput[T comparable] struct {
	Slice  []T
	Map    map[T]struct{}
	Target T
}

func makeRunInput[T comparable](sliceLen int, init func() T) RunInput[T] {
	var slice []T
	for range sliceLen {
		slice = append(slice, init())
	}

	target := slice[mathRand.IntN(sliceLen)]

	mapping := make(map[T]struct{})
	for _, v := range slice {
		mapping[v] = struct{}{}
	}

	return RunInput[T]{
		Slice:  slice,
		Map:    mapping,
		Target: target,
	}
}

func makeRunInputs[T comparable](tests int, sliceLen int, init func() T) []RunInput[T] {
	var inputs []RunInput[T]
	for range tests {
		inputs = append(inputs, makeRunInput(sliceLen, init))
	}

	return inputs
}

func BenchmarkContains(b *testing.B) {
	b.StopTimer()

	inputs := makeRunInputs(10_000, 20, rand.Text)

	b.Run("slice", func(b *testing.B) {
		b.StopTimer()
		for _, input := range inputs {
			b.StartTimer()
			ok := slices.Contains(input.Slice, input.Target)
			b.StopTimer()

			require.True(b, ok)
		}
	})

	b.Run("map", func(b *testing.B) {
		b.StopTimer()
		for _, input := range inputs {
			b.StartTimer()
			_, ok := input.Map[input.Target]
			b.StopTimer()

			require.True(b, ok)
		}
	})
}
