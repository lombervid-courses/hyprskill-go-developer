package main

import (
	"fmt"
	"time"
)

/******************************* After *******************************/

func performLongOperation(resultChan chan<- string) {
	// Simulate a long operation that takes 3 seconds.
	time.Sleep(3 * time.Second)

	// Send the result to the channel.
	resultChan <- "Operation completed successfully"
}

func timeAfter() {
	resultChan := make(chan string)

	// Start the long operation in a goroutine.
	go performLongOperation(resultChan)

	// Use select to wait for the result or timeout.
	select {
	case result := <-resultChan:
		fmt.Println(result)
	case <-time.After(2 * time.Second):
		fmt.Println("Operation timed out")
	}
}

/******************************* Timer *******************************/

func taskA() {
	fmt.Println("Task A executed at", time.Now())
}

func taskB() {
	fmt.Println("Task B executed at", time.Now())
}

func timeTimer() {
	timerA := time.NewTimer(3 * time.Second)
	timerB := time.NewTimer(5 * time.Second)

	// defer timerA.Stop()
	// defer timerB.Stop()

	for {
		select {
		case <-timerA.C:
			taskA()
			// Reset the timer for Task A to run again in 3 seconds.
			timerA.Reset(3 * time.Second)

		case <-timerB.C:
			taskB()
			// Reset the timer for Task B to run again in 5 seconds.
			timerB.Reset(5 * time.Second)
		}
	}
}

func timeTimer2() {
	timer := time.NewTimer(time.Second)

	go func() {
		<-timer.C
		fmt.Println("Timer expired")
	}()

	stop := timer.Stop()
	if stop {
		fmt.Println("Timer stopped")
	}
}

/******************************* Ticker *******************************/

func timeTicker() {
	fmt.Println("Start")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			fmt.Println("Tick")
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Println("End")
}

func main() {
	fmt.Println("Time scheduling")

	// timeAfter()
	// timeTimer()
	// timeTimer2()
	timeTicker()
}
