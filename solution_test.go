package ohmygo

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

func TestPrime20(t *testing.T) {
	result := prime(20)
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nprime in 1..10 is 2, 3, 5, 7, 11, 13, 17, 19\nbut you've got %#v", result)
	}
}

func TestPrime25(t *testing.T) {
	result := prime(25)
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nprime in 1..10 is 2, 3, 5, 7, 11, 13, 17, 19, 23\nbut you've got %#v", result)
	}
}

func TestPrime100(t *testing.T) {
	result := prime(100)
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nprime in 1..10 is 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97\nbut you've got %#v", result)
	}
}
