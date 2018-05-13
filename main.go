package main

import "fmt"

// http://keshavabharadwaj.com/2016/03/31/closure_golang/
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
