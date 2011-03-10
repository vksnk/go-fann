package main

import (
	"fann"
)

func main() {
	const numLayers = 3
	const desiredError = 0.001
	const maxEpochs = 500000
	const epochsBetweenReports = 1000

	ann := fann.CreateStandart(numLayers, []uint32{2, 3, 1})
	ann.TrainOnFile("xor.data", maxEpochs, epochsBetweenReports, desiredError)
	ann.Foo()
}
