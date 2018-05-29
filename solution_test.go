package oh_my_go

import (
	"testing"
)

func TestFizzBuzz1(t *testing.T) {
	result := fizzbuzz(1)
	if result != "1" {
		t.Error("1 should be 1 but got", result)
	}
}
