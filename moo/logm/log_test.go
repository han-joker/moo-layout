package logm

import (
	"testing"
)

//func TestInstance(t *testing.T) {
//	i1 := Inst()
//	i2 := Inst()
//	if i1 != i2 {
//		t.Error("no singleton")
//	}
//}

//func TestLog_Info(t *testing.T) {
//	i := Inst(Opt{
//		Fmt: JSON,
//		Caller: false,
//		Mode: FILE,
//	})
//	i.Info("some message")
//}

//func BenchmarkInstance(b *testing.B) {
//	for i := 0; i < b.N; i ++ {
//		Inst()
//	}
//}

func BenchmarkLog_Info(b *testing.B) {
	ins := Inst(Opt{
		Mode: FILE,
		Path: "./logs",
		Filename: "app-%y-%m",
	})
	for i := 0; i < b.N; i ++ {
		ins.Info("some message")
	}
}

