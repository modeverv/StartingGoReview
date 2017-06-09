package foo

import (
	"testing"
)

func TestIsOne(t *testing.T) {
	n := 1
	b := IsOne(1)
	if b != true {
		t.Errorf("%d is not one", n)
	}
}

