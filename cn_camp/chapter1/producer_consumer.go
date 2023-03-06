package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go producer(ch)
	go consumer(ch)
	time.Sleep(11 * time.Second)
}

func consumer(ch chan int) {
	for {
		d, ok := <-ch
		if ok {
			fmt.Println("consume value:", d)
		} else {
			fmt.Println("channel closed, consumer quit")
			return
		}

	}
}

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
