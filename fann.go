package fann

/*
#include <fann.h>

typedef unsigned int* uint_ptr;

const unsigned int* go2uintArray(unsigned int* arr, int n) {
	return malloc(sizeof(unsigned int) * n);
}

*/
import "C"

type TrainData struct {
	object *C.struct_fann_train_data
}

func ReadFromFile(fn string) {
}

type Ann struct {
	object *C.struct_fann
}

func CreateStandart(numLayers uint, layers []uint32) (*Ann) {
	var ann Ann
	ann.object = C.fann_create_standard_array(C.uint(numLayers), (*C.uint)(&layers[0]))
	return &ann
}
/*
func CreateSparse(concentration float32, num_layers uint, layers []int) {
}
*/
