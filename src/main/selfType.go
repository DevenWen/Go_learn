package main

import "fmt"

func main() {
	type (
		user struct {
			name string
			age uint8
		}
		/*Note 函数类型*/
		event func(string) bool
	)

	u := user{"Tom", 20}
	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	f("abc")
}
