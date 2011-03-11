package main

import (
	"fann"
	"fmt"
)

func main() {
	input := []fann.FannType{-1.0, 1.0}

	ann := fann.CreateFromFile("xor_float.net")

	output := ann.Run(input)

	fmt.Printf("xor test (%f,%f) -> %f\n", input[0], input[1], output[0])
	ann.Destroy()
}
