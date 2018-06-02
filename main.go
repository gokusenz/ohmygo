package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	go func() {
		s := <-sig
		fmt.Println("bye :)", s, s.Signal())
		os.Exit(0)
	}()

	go func() {
		for {
			fmt.Println()
			<-time.After(time.Second)
		}
	}()

	c := time.Tick(100 * time.Millisecond)

	for range c {
		fmt.Print(".")
	}
}
