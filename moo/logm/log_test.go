package logm

import (
	"os"
	"testing"
)

func TestInstance(t *testing.T) {
	i1 := Inst()
	i2 := Inst()
	if i1 != i2 {
		t.Error("no singleton")
	}
}

func TestLog_Info(t *testing.T) {
	i := Inst(Opt{
		Fmt: "json",
		Caller: false,
		Out: os.Stderr,
	})
	i.Info("some message")
}

func BenchmarkInstance(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		Inst()
	}
}
