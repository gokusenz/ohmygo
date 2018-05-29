package oh_my_go

import (
	"reflect"
	"testing"
)

func TestPrime10(t *testing.T) {
	result := prime(10)
	expected := []int{2, 3, 5, 7}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("prime in 1..10 is 2,3,5,7 but you've got %#v", result)
	}
}
