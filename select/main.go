package main

import (
	"fmt"
	"time"
)

// func main() {
// 	chan1 := make(chan int)
// 	chan2 := make(chan int)

// 	var data int
// 	go func() {
// 		for {
// 			fmt.Scan(&data)
// 			// write
// 			select {
// 			case chan1 <- data:
// 			case chan2 <- data:
// 			}
// 		}
// 	}()

// 	for range [3]struct{}{} {
// 		// read
// 		select {
// 		case data := <-chan1:
// 			fmt.Println("Data from chan1:", data)
// 		case data := <-chan2:
// 			fmt.Println("Data from chan2:", data)
// 		}
// 	}
// }

// Input:
// 200 403 505
// Output:
// Data from chan1: 200
// Data from chan2: 403
// Data from chan2: 505

/**************************** Default Branch ****************************/

// func main() {
// 	JackTasks := make(chan int, 1)
// 	RoseTasks := make(chan int, 1)

// 	var taskID int

// 	for {
// 		fmt.Scan(&taskID)

// 		select {
// 		case JackTasks <- taskID:
// 			fmt.Printf("The task (%d) is assigned to Jack\n", taskID)
// 		case RoseTasks <- taskID:
// 			fmt.Printf("The task (%d) is assigned to Rose\n", taskID)
// 		default:
// 			fmt.Println("There are no available employees")
// 			return
// 		}
// 	}
// }

// Input:  32
// Output: The task (32) is assigned to Jack
// Input:  64
// Output: The task (64) is assigned to Rose
// Input:  128
// Output: There are no available employees

/**************************** Timeout ****************************/

// func main() {
// 	JackTasks := make(chan int, 1)

// 	var taskID int
// 	fmt.Scan(&taskID)

// 	JackTasks <- taskID                     // assign the task to Jack
// 	timer := time.NewTimer(time.Second * 3) // set time to complete the task
// 	defer timer.Stop()

// 	for {
// 		select {
// 		case <-timer.C:
// 			fmt.Printf("Jack has finished task (%d)", <-JackTasks)
// 			return
// 		default:
// 			fmt.Println("Jack is working")
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

/**************************** Block Forever ****************************/

func main() {
	go func() {
		for {
			fmt.Println("The task is executed every 2 seconds")
			time.Sleep(2 * time.Second)
		}
	}()
	go func() {
		for {
			fmt.Println("The task is executed every 3 seconds")
			time.Sleep(3 * time.Second)
		}
	}()

	select {}
}

// Output:
// The task is executed every 3 seconds
// The task is executed every 2 seconds
// The task is executed every 2 seconds
// The task is executed every 3 seconds
// The task is executed every 2 seconds
// ...
