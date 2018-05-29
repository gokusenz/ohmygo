package oh_my_go

import (
	"testing"
)

func TestFactorial3(t *testing.T) {
	result := factorial(3)
	if result != 6 {
		t.Error("factorial of 3 is 6 but you've got", result)
	}
}

func TestFactorial4(t *testing.T) {
	result := factorial(4)
	if result != 24 {
		t.Error("factorial of 4 is 24 but you've got", result)
	}
}

func TestFactorial5(t *testing.T) {
	result := factorial(5)
	if result != 120 {
		t.Error("factorial of 5 is 120 but you've got", result)
	}
}

func TestFactorial6(t *testing.T) {
	result := factorial(6)
	if result != 720 {
		t.Error("factorial of 6 is 720 but you've got", result)
	}
}
