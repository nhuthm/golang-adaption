package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}

	//Name: syntax - subset of field
	v2 = Vertex{Y: 2}
	v3 = Vertex{}

	//& - returns a pointer to the struct value
	p = &Vertex{1, 2}	
)

func main() {
	fmt.Println(v1, p, v2, v3)
}