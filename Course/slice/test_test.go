package main

import "testing"

func TestAppendSilce(t *testing.T) {
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

func BenchmarkAppendSilce(b *testing.B) {
	for i:=0;i<b.N;i++ {
		AppendSilce()
	}
}