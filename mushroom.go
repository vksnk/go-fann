package main

import (
	"fann"
	"fmt"
)

func main() {
	const numLayers = 3
	fmt.Println("Creating network.")
	ann := fann.CreateStandart(numLayers, []uint32{2, 3, 1})
	ann.Destroy()
}
