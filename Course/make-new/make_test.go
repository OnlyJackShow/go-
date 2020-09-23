package main

import "testing"

func TestMakeFunc(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func BenchmarkMakeFunc(b *testing.B) {
	for i:=0;i<b.N;i++ {
		MakeFunc()
	}
}

func BenchmarkMakeMap(b *testing.B) {
	for i:=0;i<b.N;i++ {
		MakeMap()
	}
}

func BenchmarkMakeChan(b *testing.B) {
	MakeChan()
}

func BenchmarkMakeCap(b *testing.B) {
	MakeCap()
}

func BenchmarkTesttest(b *testing.B) {
	for i:=0;i<b.N;i++ {
		Testtest()
	}
}
