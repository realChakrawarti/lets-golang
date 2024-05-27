package main

import (
	"fmt"
	"time"
)

func greet(message string, greetDoneChan chan bool) {
	fmt.Println(message)
	greetDoneChan <- true
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

func slowGreet(message string, greetDoneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(message)
	greetDoneChan <- true
	// Closes the channel when its not needed anymore
	// Sending on closed channel throws -> panic: send on closed channel
	close(greetDoneChan)
}

func main() {
	iterateDone := make(chan bool)
	greetDone := make(chan bool)
	// Re-use the same channel but this order doesn't matter
	go greet("AI: Hello, how are you?", greetDone)
	go greet("Me: Who do you think you are!", greetDone)
	// iterate(18446744073709551615)
	go iterate(10000000000, iterateDone)
	go slowGreet("Me: Great!", greetDone)
	<-greetDone
	<-iterateDone
}
