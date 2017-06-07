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

func swap(p *Point){
	x,y := p.y,p.x
	p.x = x
	p.y = y
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
	p := Point{1,2}
	f.Println(p)
	swap(&p)
	f.Println(p)
}
