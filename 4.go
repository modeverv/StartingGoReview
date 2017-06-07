package main

import f "fmt"

type Feed struct {
	Name   string
	Amount uint
}

type Animal struct {
	Name string
	Feed Feed
}

type Point struct {
	x,y int
}

func main() {
	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}
	f.Println(a)
}
