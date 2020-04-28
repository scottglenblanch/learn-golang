package main

import "fmt"

const (
	id1 = iota + 1
	id2
	id3
	id4
	id5
)

func main() {
	fmt.Println("Lesson 4")
	fmt.Println(id1, id2, id3, id4, id5)
}

