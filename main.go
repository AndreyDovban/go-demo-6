package main

import "fmt"

func main() {
	x := 10
	y := double(x)

	fmt.Println(y)
}

func double(x int) int {
	return x * 2
}
