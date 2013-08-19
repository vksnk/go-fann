package main

import (
	"github.com/white-pony/go-fann"
	"fmt"
)
func main() {
	const num_layers = 3
	const num_neurons_hidden = 96
	const desired_error = 0.001

	train_data := fann.ReadTrainFromFile("../../datasets/robot.train")
	test_data := fann.ReadTrainFromFile("../../datasets/robot.test")

	var momentum float32
	for momentum = 0.0; momentum < 0.7; momentum += 0.1 {
		fmt.Printf("============= momentum = %f =============\n", momentum)

		ann := fann.CreateStandard(num_layers, []uint32{train_data.GetNumInput(), num_neurons_hidden, train_data.GetNumOutput()})
		ann.SetTrainingAlgorithm(fann.TRAIN_INCREMENTAL)
		ann.SetLearningMomentum(momentum)
		ann.TrainOnData(train_data, 2000, 500, desired_error)

		fmt.Printf("MSE error on train data: %f\n", ann.TestData(train_data))
		fmt.Printf("MSE error on test data : %f\n", ann.TestData(test_data))

		ann.Destroy()
	}

	train_data.Destroy()
	test_data.Destroy()
}
