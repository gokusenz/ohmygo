package ohmygo

import (
	"reflect"
	"testing"
)

func TestCountChar1(t *testing.T) {
	m := each("I am learning Go")
	expected := map[string]int{
		"I":        1,
		"am":       2,
		"learning": 8,
		"Go":       2,
	}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("\n%#v\nis expected but got\n%#v\n", expected, m)
	}
}

func TestCountChar2(t *testing.T) {
	m := each("The quick brown fox jumped over the lazy dog")
	expected := map[string]int{
		"The":    3,
		"quick":  5,
		"brown":  5,
		"fox":    3,
		"jumped": 6,
		"over":   4,
		"the":    3,
		"lazy":   4,
		"dog.":   3,
	}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("\n%#v\nis expected but got\n%#v\n", expected, m)
	}
}
