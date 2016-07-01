package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d:%d\n", n, i)
	}
}

func main() {
	go f(0)
	time.Sleep(1 * time.Second)
}