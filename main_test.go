package main

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

var slice []int
var m = make(map[int]struct{})

func genericContains(n int) bool {
	return slices.Contains(slice, n)
}

func mapCheck(n int) bool {
	_, ok := m[n]
	return ok
}

func sliceCheck(n int) bool {
	for _, v := range slice {
		if v == n {
			return true
		}
	}
	return false
}

func setup(sliceSize int) {
	for i := 0; i < sliceSize; i++ {
		slice = append(slice, i)
	}
}

func constructMap() {
	for _, value := range slice {
		m[value] = struct{}{}
	}
}

func BenchmarkContainsCheck(b *testing.B) {
	var table = []struct {
		sliceSize int
		inProb    float64
	}{
		{sliceSize: 10, inProb: 0.1},
		{sliceSize: 10, inProb: 0.5},
		{sliceSize: 100, inProb: 0.1},
		{sliceSize: 100, inProb: 0.5},
		{sliceSize: 1000, inProb: 0.1},
		{sliceSize: 1000, inProb: 0.5},
	}

	for _, v := range table {
		setup(v.sliceSize)
		inCounts := int(float64(b.N) * v.inProb)
		inNumber := v.sliceSize / 2
		outNumber := v.sliceSize + 1
		b.Run(fmt.Sprintf("construct map : size %d", v.sliceSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				constructMap()
			}
		})
		b.Run("map", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if i < inCounts{
					mapCheck(inNumber)
					continue
				}
				mapCheck(outNumber)
			}
		})
		b.Run("slice", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if i < inCounts{
					sliceCheck(inNumber)
					continue
				}
				sliceCheck(outNumber)
			}
		})
		b.Run("generic slice contains", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if i < inCounts{
					genericContains(inNumber)
					continue
				}
				genericContains(outNumber)
			}
		})
		slice = []int{}
	}
}
