package main

import f "fmt"

func reseiver(ch chan int) {
	for {
		i := <-ch
		f.Println(i)
	}
}

func main() {
	ch := make(chan int)

	go reseiver(ch)

	i := 0
	for i < 10 {
		ch <- i
		i++
	}
}
