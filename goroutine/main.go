package main

import (
	"fmt"
	"time"
)

func subRoutine(isGoroutine bool, seq int) {
	var message string
	message = "normal"
	if isGoroutine {
		message = "goroutine"
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("[%d]%s i : %d\n", seq, message, i)
	}
}

func main() {
	subRoutine(false, 0)

	go subRoutine(true, 1)
	go subRoutine(true, 2)

	fmt.Println("Sleep Thread for 1 seconds")
	time.Sleep(time.Second)
}
