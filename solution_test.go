package ohmygo

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	answer := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	f := fibonacci()
	for i := 0; i < 10; i++ {
		if x := f(); answer[i] != x {
			t.Errorf("answer is %d but got %d", answer[i], x)
		}
	}
}
