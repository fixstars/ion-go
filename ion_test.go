package ion

import (
	"testing"
)

func Test_CFFI(t *testing.T) {
	if CFFI() != 0 {
		t.Errorf("X")
	}
}
