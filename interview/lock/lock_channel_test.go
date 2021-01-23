package main

import "testing"

func BenchmarkUseMutex(b *testing.B) {
	for n := 0; n < 1000; n++ {
		UseMutex()
	}
}
func BenchmarkUseChan(b *testing.B) {
	for n := 0; n < 1000; n++ {
		UseChan()
	}
}

//go test -v -bench=.
