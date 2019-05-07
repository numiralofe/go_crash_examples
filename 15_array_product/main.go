package main

import "fmt"

func main() {
	adding, product := 2, 5
	for _, x := range []int{2, 5} {
		adding += x
		product *= x
	}
	fmt.Println(adding, product)
}
