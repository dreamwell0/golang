package main

import "fmt"

func main() {
	x := 3
	y := f(x)
	fmt.Println(y)
}

// f(x) = x + 2という式を意味する関数
func f(x int) int {
	y := x + 2
	return y
}
