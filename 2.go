package main

// 配列！！！！！

import "fmt"

func main() {
	//n := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//fmt.Println(n)

	//s := make([]int, 8)
	//fmt.Println(len(s))
	//fmt.Println(s)
	//s = append(s, 8)
	//fmt.Println(s)

	//配列の 構えは
	n := 10
	s := make([]int, n)
	// 拡張
	x := 11
	s = append(s, x)
	fmt.Println(len(s))
	//回すのは
	for i, v := range s {
		fmt.Printf("%d - %d\n", i, v)
	}

}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
