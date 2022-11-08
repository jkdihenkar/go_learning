package main

import (
	"fmt"
	"time"
)

/*
	You can stop a goroutine by sending a signal channel
*/
func main() {
	quit := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Printf("Running routine - iter %d\n", i)
				i = i + 1
				time.Sleep(3 * time.Second)
			}
		}
	}()
	// something happens ...
	time.Sleep(9 * time.Second)
	fmt.Println("Ticker Stop - quit <- true")
	quit <- true
}

/*
	Running routine - iter 1
	Running routine - iter 2
	Running routine - iter 3
	Ticker Stop - quit <- true
*/
