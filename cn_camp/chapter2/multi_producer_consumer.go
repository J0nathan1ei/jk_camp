package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 100)
	var consumerWg, producerWg sync.WaitGroup

	// 10个生产者
	for i := 0; i < 10; i++ {
		producerWg.Add(1)
		go producer(ch, i, &producerWg)
	}
	// 10个消费者
	for i := 0; i < 10; i++ {
		consumerWg.Add(1)
		go consumer(ch, &consumerWg)
	}

	producerWg.Wait()
	close(ch)
	consumerWg.Wait()
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range ch {
		fmt.Println("consume value:", d)
	}
	fmt.Println("channel closed, consumer quit")

}

func producer(ch chan int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i * n
		time.Sleep(1 * time.Second)
	}
}
