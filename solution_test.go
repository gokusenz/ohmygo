package ohmygo

import (
	"fmt"
	"testing"
)

func TestRacing(t *testing.T) {
	data := [][]float64{
		{170, 70, 24.221453287197235},
		{150, 65, 28.88888888888889},
		{165, 80, 29.384756657483933},
		{173, 75, 25.05930702662969},
	}

	ch := make(chan float64)

	for _, person := range data {
		go bmi(person[0], person[1], ch)
	}

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
}
