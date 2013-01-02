package fann



import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	input := []FannType{-1.0, 1.0}

	// can't find file: xor_float.net in the latest FANN distribution
	ann := CreateFromFile("xor_float.net")
	
	output := ann.Run(input)

	fmt.Printf("xor test (%f,%f) -> %f\n", input[0], input[1], output[0])
	ann.Destroy()
}
