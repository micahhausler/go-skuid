package main

import "fmt"
import "time"

func worker(c chan bool, id int, finish bool) {
	fmt.Printf("working on %d\n", id)
	if finish {
		time.Sleep(time.Second * 2)
		fmt.Println("done")
		c <- true
	}
}

func main() {
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go worker(c, i, false)
	}
	go worker(c, 99, true)
	fmt.Println("sent all workers")
	// Block until we receive a notification from the worker on the channel.
	<-c
}
