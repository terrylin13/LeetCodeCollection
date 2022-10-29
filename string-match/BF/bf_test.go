package bf

import (
	"fmt"
	"testing"
)

func TestBF(t *testing.T) {
	s := "testing"
	target := "test"
	fmt.Println(BF(s, target))
}

func TestBFAll(t *testing.T) {
	s := "testing test"
	target := "test"
	fmt.Println(BFAll(s, target))
}

func TestReplace(t *testing.T) {
	s := "testing"
	target := "test"
	want := "yes"
	t.Log(Replace(s, target, want))
}
