package maptypes

import (
	"fmt"
	"strconv"
	"testing"
)

type (
	payload struct {
		A, B, C int
		D, E, F string
		G       [10]byte
		H       struct {
			A int
			B string
		}
	}
)

func BenchmarkMapUint(b *testing.B) {
	set := []int{1000, 10000, 100000, 1000000}
	for i := range set {
		max := set[i]
		b.Run(fmt.Sprintf("%d", max), func(b *testing.B) {
			data := make(map[uint]payload, max)
			for i := 1; i <= max; i++ {
				data[uint(i)] = payload{}
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = data[uint(i%max)]
			}
		})
	}
}

func BenchmarkMapString(b *testing.B) {
	set := []int{1000, 10000, 100000, 1000000}
	for i := range set {
		max := set[i]
		b.Run(fmt.Sprintf("%d", max), func(b *testing.B) {
			data := make(map[string]payload, max)
			for i := 1; i <= max; i++ {
				data[strconv.Itoa(i)] = payload{}
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = data[strconv.Itoa(i%max)]
			}
		})
	}
}

func BenchmarkMapStruct(b *testing.B) {
	set := []int{1000, 10000, 100000, 1000000}
	for i := range set {
		max := set[i]
		b.Run(fmt.Sprintf("%d", max), func(b *testing.B) {
			type key struct{ A, B int }
			data := make(map[key]payload, max)
			for i := 1; i <= max; i++ {
				data[key{i, i}] = payload{}
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = data[key{i, i}]
			}
		})
	}

}
