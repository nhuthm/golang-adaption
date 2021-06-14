package main

import "fmt"

func returnValue(sum int) (x, y int) {
	x = sum + 1
	y = x + sum
	return
}

func main() {
	fmt.Println(returnValue(23))
}
