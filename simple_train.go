package main

import (
	"fann"
)

func main() {
	const numLayers = 3
	ann := fann.CreateStandart(numLayers, []uint32{2, 3, 1})
	ann.Foo()
}
