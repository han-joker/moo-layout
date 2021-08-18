package confm

import (
	"testing"
)

func TestInstance(t *testing.T) {
	m1 := Instance()
	m2 := Instance()
	if m1 != m2 {
		t.Error("no singleton")
	}
}

func BenchmarkInstance(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		Instance()
	}
}