package main

import(
	f "fmt"
)

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{"深刻なエラーが発生しました", 1234}
}

func main() {
	err := RaiseError()
	// 型アサーション
	e, ok := err.(*MyError)
	if ok {
		f.Println(e.ErrCode)
	}
}
