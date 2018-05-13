package main

import "fmt"

func outer(name string) {
	text := "Hello " + name

	inner := func() {
		fmt.Println(text)
	}

	inner()
}

func main() {
	outer("pallat")
}
