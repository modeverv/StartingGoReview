package main

import(
	f "fmt"
	"encoding/json"
)

type Poin struct {
	X int `json:"x"`
  Y int `json:"y"`
}

type Poin2 struct {
	X int
	Y int
}

func main(){
	p := Poin{1,2}
	p2 := Poin2{1,2}
	f.Println(p)
	bs,_ := json.Marshal(p)
	bs2,_ := json.Marshal(p2)
	// JSONで書き出すためにはフィールドが可視でないとだめ。
	f.Println(string(bs))
	f.Println(string(bs2))

}
