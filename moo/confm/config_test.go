package confm

import (
	"testing"
)

func TestInstance(t *testing.T) {
	m1 := Get()
	m2 := Get()
	if m1 != m2 {
		t.Error("no singleton")
	}
}

func BenchmarkInstance(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		//New()
		Get()
	}
}