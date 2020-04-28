package main

import "fmt"

func main() {

	varRegular := "This is a string"
	varPointer := &varRegular

	fmt.Println("Lesson 3")
	fmt.Println("varRegular", varRegular)
	fmt.Println("varPointer memory", varPointer, "varPointer value", *varPointer)
}

