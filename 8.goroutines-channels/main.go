package main

import (
	"fmt"
	"time"
)

func sayHello(message string) {
	fmt.Println(message)
}

func main() {
	// goroutine
	go sayHello("Hello")        // Run as a goroutine
	time.Sleep(7 * time.Second) // Wait to let goroutine complete

	// channel
	messages := make(chan string)

	go func() {
		messages <- "Hello from channel"
	}()

	fmt.Println(<-messages)
}
