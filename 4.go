package main

import (
	f "fmt"
	"math"
)

type Feed struct {
	Name   string
	Amount uint
}

type Animal struct {
	Name string
	Feed Feed
}

type Point struct {
	x, y int
}

func swap(p *Point) {
	x, y := p.y, p.x
	p.x = x
	p.y = y
}

// 構造体の「メソッド」
func (p *Point) Distance(dp *Point) float64 {
	x, y := p.x-dp.x, p.y-dp.y
	return math.Sqrt(float64(x*x + y*y))
}

// 頻出のコンストラクタっぽいの
func NewPoint(x int, y int) *Point {
	p := new(Point)
	p.x, p.y = x, y
	return p
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
	p := Point{1, 2}
	f.Println(p)
	swap(&p)
	f.Println(p)

	p1 := Point{1, 2}
	d := p1.Distance(&Point{4, 3})
	f.Println(d)

	p3 := NewPoint(3, 3)
	f.Println(p3)

	// マップの初期化で構造体は楽できる例
	mm := map[int]Point{
		1: {3, 2},
		2: {4, 6},
	}
	f.Println(mm)
	mp,ok := mm[1]
	if ok {
		f.Println(mp.x)
	}

}
