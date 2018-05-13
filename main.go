package main

import "fmt"

func outer(name string) func() {
	text := "Hello " + name

	inner := func() {
		fmt.Println(text)
	}

	return inner
}

func main() {
	fn := outer("pallat")

	fn()
}
