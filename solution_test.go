package oh_my_go

import (
	"testing"
)

func TestMan(t *testing.T) {
	man := newMan()
	if !man.canWalk() {
		t.Errorf("man should walkable!")
	}
}

func TestDog(t *testing.T) {
	micky := newDog()
	if !micky.canWalk() {
		t.Errorf("dog should walkable!")
	}
}
