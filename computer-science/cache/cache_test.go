package main

import (
	"testing"
)

const arraySize int = 10000

func twoDimensionalArray() [][]int {
	arr := make([][]int, arraySize)
	for i := 0; i < arraySize; i++ {
		arr[i] = make([]int, arraySize)
		for j := 0; j < arraySize; j++ {
			arr[i][j] = 0
		}
	}
	return arr
}

func BenchmarkSpatialLocality(b *testing.B) {
	// 바로 다음 메모리에 접근
	b.Run("Use Spatial Locality", func(b *testing.B) {
		arr := twoDimensionalArray()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < arraySize; i++ {
				arr[0][i] += 1
			}
		}
	})

	// 멀리있는 메모리에 접근
	b.Run("Not Use Spatial Locality", func(b *testing.B) {
		arr := twoDimensionalArray()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < arraySize; i++ {
				arr[i][0] += 1
			}
		}
	})
	/*
		goos: linux
		goarch: amd64
		pkg: github.com/ice-coldbell/study/computer-science/cache
		cpu: 13th Gen Intel(R) Core(TM) i5-13600K
		BenchmarkSpatialLocality/Use_Spatial_Locality-20         	  118550	     10092 ns/op
		BenchmarkSpatialLocality/Not_Use_Spatial_Locality-20     	    9709	    139221 ns/op
	*/
}

func BenchmarkTemporalLocality(b *testing.B) {
	// 10개의 데이터에만 접근
	b.Run("Use Temporal Locality x10", func(b *testing.B) {
		arr := twoDimensionalArray()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < arraySize; i++ {
				arr[i%10][0] += 1
			}
		}
	})

	// 새로운 데이터에만 접근
	b.Run("Not Use Temporal Locality", func(b *testing.B) {
		arr := twoDimensionalArray()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < arraySize; i++ {
				arr[i][0] += 1
			}
		}
	})
	/*
		goos: linux
		goarch: amd64
		pkg: github.com/ice-coldbell/study/computer-science/cache
		cpu: 13th Gen Intel(R) Core(TM) i5-13600K
		BenchmarkTemporalLocality/Use_Temporal_Locality_x8-20         	   52437	     20697 ns/op
		BenchmarkTemporalLocality/Not_Use_Temporal_Locality-20        	    9484	    141967 ns/op
	*/
}
