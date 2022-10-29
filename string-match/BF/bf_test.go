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
