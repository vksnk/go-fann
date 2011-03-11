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

	testData := fann.ReadTrainFromFile("datasets/mushroom.test")

	ann.ResetMSE()
	/*
	for(i = 0; i < fann_length_train_data(test_data); i++)
	{
		fann_test(ann, test_data->input[i], test_data->output[i]);
	}
	*/
	fmt.Printf("MSE error on test data: %f\n", ann.GetMSE())

	fmt.Println("Saving network.");

	ann.Save("mushroom_float.net");

	fmt.Println("Cleaning up.");

	trainData.Destroy()
	testData.Destroy()
	ann.Destroy()
}
