package main

import (
	"testing"
)

const all = 0xffffffffffffffff
const half = 0x0f0f0f0f0f0f0f0f
const sparse = 0x1001000001000001
const none = 0x0

type Test struct {
	w   uint64
	exp int
}

var testData = [4]Test{{all, 64}, {half, 32}, {sparse, 4}, {none, 0}}

func TestSkipZeros(t *testing.T) {
	for _, test := range testData {
		act := SkipZeros(test.w)
		if act != test.exp {
			t.Errorf("SkipZeros(%d) = %d; want %d", test.w, act, test.exp)
		}
	}
}

func BenchmarkSkipZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testData {
			SkipZeros(test.w)
		}
	}
}

func TestLogBitSum(t *testing.T) {
	for _, test := range testData {
		act := LogBitSum(test.w)
		if act != test.exp {
			t.Errorf("SkipZeros(%d) = %d; want %d", test.w, act, test.exp)
		}
	}
}

func BenchmarkLogBitSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testData {
			LogBitSum(test.w)
		}
	}
}
