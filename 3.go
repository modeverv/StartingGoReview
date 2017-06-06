package main

import (
	"fmt"
)

func receiver(ch <-chan int) {
	for {
		i, ok := <-ch
		if ok {
			fmt.Println(i)
		}
	}
}

func main() {
	ch := make(chan int)
	go receiver(ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)
}
