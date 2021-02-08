package moduleb

import (
	"testing"
)

func TestToCaps(t *testing.T) {
	s := ToCaps("heLLo")
	if s != "HELLO" {
		t.Fail()
	}
}