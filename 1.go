package main

import (
	f "fmt"
)

func main() {
	m := map[string]string{
		"a": "hoge",
		"b": "fuga",
	}
	s, ok := m["a"]

	if ok {
		f.Println(s)
	}

	for k, v := range m {
		f.Printf("%s - %s\n", k, v)
	}
}
