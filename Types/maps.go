package main

//syntax: make(map[type of key]type of value)

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// var personSalary map[string]int
	// if personSalary == nil {
	// 	fmt.Println("Map is nil right now")
	// 	personSalary = make(map[string]int)
	// }

	// personSalary := make(map[string]int)
	// personSalary["nhut"] = 12000
	// personSalary["huy"] = 2000
	// personSalary["quynh"] = 200000
	// fmt.Println("Person Salary map", personSalary)

	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.5, 80.5,
	}
	fmt.Println(m["Bell Labs"])
}
