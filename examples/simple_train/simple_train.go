package main

import (
	"fann"
)

func main() {
	const numLayers = 3
	const desiredError = 0.00001
	const maxEpochs = 500000
	const epochsBetweenReports = 1000

	ann := fann.CreateStandard(numLayers, []uint32{2, 3, 1})
	ann.SetActivationFunctionHidden(fann.SIGMOID_SYMMETRIC)
	ann.SetActivationFunctionOutput(fann.SIGMOID_SYMMETRIC)
	ann.TrainOnFile("../../datasets/xor.data", maxEpochs, epochsBetweenReports, desiredError)
	ann.Save("xor_float.net")
	ann.Destroy()
}
