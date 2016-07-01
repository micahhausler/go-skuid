package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d:%d\n", n, i)
	}
}

func main() {
	go f(0)
	go f(1)
	time.Sleep(1 * time.Second)
}
