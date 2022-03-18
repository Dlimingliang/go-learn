package main

import (
	"testing"
)

//普通测试
func TestAdd(t *testing.T) {
	test := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{1, 2, 3},
	}

	for _, tt := range test {
		if actual := add(tt.a, tt.b); actual != tt.c {
			t.Errorf("params %d + %d, expect %d, but actual %d", tt.a, tt.b, tt.c, actual)
		}
	}
}

//性能测试
func BenchmarkAdd(b *testing.B) {
	a, d, c := 10000, 10000, 20000
	for i := 0; i < b.N; i++ {
		if actual := add(a, d); actual != c {
			b.Errorf("params %d + %d, expect %d, but actual %d", a, d, c, actual)
		}
	}
}
