package main

import "fmt"

func functions() []func() {
	// pitfall of using loop variables
	arr := []int{1, 2, 3, 4}
	result := make([]func(), 0)

	// for i := range arr {
	for i := 0; i < 3; i++ {
		fn := func() { fmt.Printf("index - %d, value - %d\n", i, arr[i]) }
		result = append(result, fn)
	}

	return result
}

type getter func() int

func pitfall() []getter {
	arr := []int{2, 3, 4, 5}
	r := make([]getter, 0)

	for _, i := range arr {
		fn := func() int { return i }
		r = append(r, fn)
	}
	return r
}

func main() {
	fns := functions()
	for f := range fns {
		fns[f]()
	}

	fnns := pitfall()

	for _, f := range fnns {
		fmt.Println(f())
	}
}
