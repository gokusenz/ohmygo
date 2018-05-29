package oh_my_go

import (
	"testing"
)

func TestIncreaseMethod1(t *testing.T) {
	num := number{n: 0}
	num.inc()
	if num.n != 1 {
		t.Errorf("it should increase from 0 to 1 but got", num.n)
	}
}
