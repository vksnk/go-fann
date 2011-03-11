package main

import (
	"fann"
	"fmt"
)

func main() {
	const numLayers = 3
	const numNeuronsHidden = 32
	const desiredError = 0.0001
	const maxEpochs = 300
	const epochsBetweenReports = 10

	fmt.Println("Creating network.")

	trainData := fann.ReadTrainFromFile("datasets/mushroom.train")
	ann := fann.CreateStandart(numLayers, []uint32{trainData.GetNumInput(), numNeuronsHidden, 1})

	fmt.Println("Training network.")
	ann.SetActivationFunctionHidden(fann.SIGMOID_SYMMETRIC_STEPWISE)
	ann.SetActivationFunctionOutput(fann.SIGMOID_SYMMETRIC)

	ann.TrainOnData(trainData, maxEpochs, epochsBetweenReports, desiredError)

	fmt.Println("Testing network")

	ann.Destroy()
}
