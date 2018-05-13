package main

import "fmt"

func counter(state int) (func() int, func()) {
	get := func() int {
		return state
	}

	inc := func() {
		state++
	}

	return get, inc
}

func main() {
	get1, inc1 := counter(100)
	get2, inc2 := counter(100)
	fmt.Println("counter1 - ", get1())
	fmt.Println("counter2 - ", get2())

	inc1()
	fmt.Println("counter1 - ", get1())
	fmt.Println("counter2- ", get2())

	inc2()
	inc2()
	fmt.Println("counter1 - ", get1())
	fmt.Println("counter2- ", get2())
}
