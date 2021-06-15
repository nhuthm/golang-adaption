package main

import "fmt"

func conditionCheck(n int) int {
	if n = n + 2; n < 3 {
		return n + 1
	} else {
		return n
	}
}

func main() {
	fmt.Println(conditionCheck(0), conditionCheck(4))
}
