package main

import (
	"fmt"
	"time"
)

// func randomInt(value chan int) {
// 	value <- rand.Intn(100)
// }

// func main() {
// 	fmt.Println("Hello, channels")

// 	var val = make(chan int)
// 	for i := 0; i < 10; i++ {
// 		go randomInt(val)
// 		fmt.Printf("%d = %d\n", i, <-val)
// 	}
// }

// doWork simulates a long computations with random execution times
// func doWork(i int, resCh chan int) {
// 	fmt.Printf("  doWork %d started\n", i)
// 	time.Sleep(time.Duration(rand.Int63n(2000)) * time.Millisecond)
// 	resCh <- i
// 	fmt.Printf("    doWork %d finished\n", i)
// }

// func main() {
// 	fmt.Println("main started")
// 	resCh := make(chan int)
// 	start := time.Now()

// 	for i := 0; i < 3; i++ {
// 		fmt.Printf("main %d started\n", i)
// 		go doWork(i, resCh)
// 	}

// 	for i := 0; i < 3; i++ {
// 		fmt.Printf("main %d finished in %.3f second\n", <-resCh, time.Since(start).Seconds())
// 	}
// 	fmt.Println("main finished")
// }

// Output:
// main started
// main 0 started
// main 1 started
// main 2 started
//   doWork 2 started
//   doWork 0 started
//   doWork 1 started
//     doWork 1 finished
// main 1 finished in 0.780 second
//     doWork 2 finished
// main 2 finished in 1.705 second
//     doWork 0 finished
// main 0 finished in 1.915 second
// main finished

/********************* Buffered Channels *********************/
// func chanReader(bufferedCh chan string) {
// 	for i := 0; i < 4; i++ {
// 		fmt.Println(<-bufferedCh)
// 	}
// }

// func main() {
// 	bufferedCh := make(chan string, 3)
// 	fmt.Println("capacity =", cap(bufferedCh))
// 	fmt.Println("length = ", len(bufferedCh))
// 	go chanReader(bufferedCh)

// 	for _, sym := range "ABCD" {
// 		bufferedCh <- string(sym)
// 		fmt.Println("length = ", len(bufferedCh))
// 	}
// }

// Output:
// capacity = 3
// length =  0
// length =  1
// length =  2
// length =  3
// A
// B
// C
// D
// length =  0

/********************* Reading and writing from channels *********************/
// func channelReader(strCh chan string) {
// 	for sym := range strCh {
// 		fmt.Print(sym)
// 	}
// }

// func main() {
// 	hello := "Hello World!"
// 	strCh := make(chan string)
// 	go channelReader(strCh)

// 	for _, sym := range hello {
// 		strCh <- string(sym)
// 	}
// 	close(strCh)

// 	// this instruction transfers control to other goroutine
// 	// it's needed for channelReader to be able to finish reading
// 	runtime.Gosched()
// }

// Output:
// Hello World!

/********************* Close channels *********************/
func main() {
	timeCh := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		timeCh <- time.Now()
		time.Sleep(time.Second)
	}
	close(timeCh)

	for i := 0; i < 5; i++ {
		if value, ok := <-timeCh; ok {
			fmt.Println(value.Format("15:04:05"), "open")
		} else {
			fmt.Println(value.Format("15:04:05"), "closed")
		}
	}
}

// Output:
// 11:57:12 open
// 11:57:13 open
// 11:57:14 open
// 00:00:00 closed
// 00:00:00 closed
