package cachem

import (
	"testing"
)

func TestNew(t *testing.T) {
	one := New()
	two := New()
	if one == two {
		t.Error("new() same instance")
	}
}

func TestGet(t *testing.T) {
	one := Get()
	two := Get()
	if one != two {
		t.Error("get not single instance")
	}
	three := Get(Option{
		Name: "three",
	})
	four := Get(Option{
		Name: "four",
	})
	if three == four {
		t.Error("get()  same instance")
	}
}
