package logm

import (
	"testing"
)
//
func TestNew(t *testing.T) {
	//one := New()
	//two := New()
	//if one == two {
	//	t.Error("new error")
	//}
	//ins := New(Option{
	//	OutMode: File,
	//	Path: "./logs",
	//})
	//ins.Info("some message")
}
//
//func TestGet(t *testing.T) {
//	one := Get()
//	two := Get()
//	if one != two {
//		t.Error("get error")
//	}
//}
//
//func BenchmarkNew(b *testing.B) {
//	for i := 0; i < b.N; i ++ {
//		New()
//	}
//}
//func BenchmarkGet(b *testing.B) {
//	for i := 0; i < b.N; i ++ {
//		Get()
//	}
//}

func BenchmarkLog_Info(b *testing.B) {
	//ins := Get(Option{
	//	OutMode: File,
	//	Path: "./logs",
	//})
	//ins := Get(Option{
	//	OutMode: FilePerHour,
	//	Path: "./logs",
	//})
	//ins := Get(Option{
	//	OutMode: FilePerWeek,
	//	Path: "./logs",
	//})

	for i := 0; i < b.N; i ++ {
		ins := Get(Option{
			OutMode: FilePerSize,
			SizeMax: 0.5*1024*1024,
			Path: "./logs",
		})
		//ins := Get(Option{
		//	OutMode: FilePerWeek,
		//	Path: "./logs",
		//})
		ins.Info("some message")
	}
}

