package main

import (
	"fann"
	"fmt"
)

func main() {
	const numLayers = 3
	const desiredError = 0.0001
	const maxEpochs = 300
	const epochsBetweenReports = 10

	fmt.Println("Creating network.")
	ann := fann.CreateStandart(numLayers, []uint32{2, 3, 1})
	trainData := fann.ReadTrainFromFile("datasets/mushroom.train")

	fmt.Println("Training network.")
	ann.SetActivationFunctionHidden(fann.SIGMOID_SYMMETRIC_STEPWISE)
	ann.SetActivationFunctionOutput(fann.SIGMOID_SYMMETRIC)

	ann.TrainOnData(trainData, maxEpochs, epochsBetweenReports, desiredError)

	fmt.Println("Testing network")

	ann.Destroy()
}
