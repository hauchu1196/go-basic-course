package main

import (
	"fmt"
	"time"
)

func main() {

	guard := make(chan struct{})

	for i := 0; i < 10000; i++ {
		go func(i int) {
			guard <- struct{}{}
			fmt.Println(i)
			<-guard
		}(i)
	}

	// deadlock
	// panic send to close channel
	// race condition

	time.Sleep(time.Hour)
}
