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
