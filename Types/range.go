package main

import "fmt"

func main() {
	a := []float64 {2, 3, 6, 13, 20, 50}
	sum := float64(0)
	for i, v := range a {//range returns both the index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of a",sum)
}