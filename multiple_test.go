package ohmygo

import "testing"

func TestMultipleReturnValues1(t *testing.T) {
	a, b := forward(1, 2)

	if a != 2 && b != 3 {
		t.Error("it should return right value and sum of them,", a, b)
	}
}

func TestMultipleReturnValues2(t *testing.T) {
	a, b := forward(2, 3)

	if a != 3 && b != 5 {
		t.Error("it should return right value and sum of them,", a, b)
	}
}

func TestMultipleReturnValues3(t *testing.T) {
	a, b := forward(3, 5)

	if a != 5 && b != 8 {
		t.Error("it should return right value and sum of them,", a, b)
	}
}
