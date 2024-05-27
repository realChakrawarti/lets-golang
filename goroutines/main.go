package main

import (
	"fmt"
	"time"
)

func greet(message string) {
	fmt.Println(message)
}

func iterate(value uint64, iterateDoneChan chan bool) {
	firstQuarter := value / 4
	half := firstQuarter * 2
	thirdQuater := firstQuarter * 3
	for i := uint64(0); i < value; i++ {
		if i == firstQuarter {
			fmt.Println("Reached 1st quarter")
		}
		if i == half {
			fmt.Println("Reached halfway")
		}
		if i == thirdQuater {
			fmt.Println("Reached 3rd quarter")
		}
		if i == value-1 {
			fmt.Println("Done iterating")
		}
	}
	iterateDoneChan <- true
}

func slowGreet(message string, slowGreetDoneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(message)
	slowGreetDoneChan <- true
}

func main() {
	greet("AI: Hello, how are you?")
	iterateDone := make(chan bool)
	slowGreetDone := make(chan bool)
	// iterate(18446744073709551615)
	go iterate(10000000000, iterateDone)
	go slowGreet("Me: Great!", slowGreetDone)
	<-slowGreetDone
	<-iterateDone
}
