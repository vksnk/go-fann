package main

import (
	"fann"
	"fmt"
)

func main() {
	const num_layers = 3
	const num_neurons_hidden = 30
	const desired_error = 0.001


	fmt.Println("Creating network.");

	train_data := fann.ReadTrainFromFile("../../datasets/robot.train")

	ann := fann.CreateStandard(num_layers, []uint32{train_data.GetNumInput(), num_neurons_hidden, train_data.GetNumOutput()})

	fmt.Println("Training network.")

	ann.SetTrainingAlgorithm(fann.TRAIN_INCREMENTAL)
	ann.SetLearningMomentum(0.4)

	ann.TrainOnData(train_data, 3000, 10, desired_error)

	fmt.Println("Testing network")

	test_data := fann.ReadTrainFromFile("../../datasets/robot.test")

	ann.ResetMSE()

	var i uint32
	for i = 0; i < test_data.Length(); i++ {
		ann.Test(test_data.GetInput(i), test_data.GetOutput(i))
	}

	fmt.Printf("MSE error on test data: %f\n", ann.GetMSE())


	fmt.Println("Saving network.");
	ann.Save("robot_float.net")
	fmt.Println("Cleaning up.")

	train_data.Destroy()
	test_data.Destroy()
	ann.Destroy()

}
