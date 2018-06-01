package ohmygo

import (
	"testing"
)

func TestFizzBuzz1(t *testing.T) {
	result := fizzbuzz(1)
	if result != "1" {
		t.Error("1 should be 1 but got", result)
	}
}

func TestFizzBuzz2(t *testing.T) {
	result := fizzbuzz(2)
	if result != "2" {
		t.Error("2 should be 2 but got", result)
	}
}

func TestFizzBuzz3(t *testing.T) {
	result := fizzbuzz(3)
	if result != "Fizz" {
		t.Error("3 should be Fizz but got", result)
	}
}

func TestFizzBuzz4(t *testing.T) {
	result := fizzbuzz(4)
	if result != "4" {
		t.Error("4 should be 4 but got", result)
	}
}

func TestFizzBuzz5(t *testing.T) {
	result := fizzbuzz(5)
	if result != "Buzz" {
		t.Error("5 should be Buzz but got", result)
	}
}
