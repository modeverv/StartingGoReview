package main

import f "fmt"

func main() {
	n := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	f.Println(n)
}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
