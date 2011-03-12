package fann
/*
#include <fann.h>
*/
import "C"
import "unsafe"

func (td *TrainData) Lenght() (uint32) {
	return uint32(C.fann_length_train_data(td.object))
}

func MergeTrainData(td1 *TrainData, td2 *TrainData) (*TrainData) {
	var td TrainData
	td.object = C.fann_merge_train_data(td1.object, td2.object)
	return &td
}

func (td* TrainData) Duplicate() (*TrainData) {
	var td_dup TrainData
	td_dup.object = C.fann_duplicate_train_data(td.object)
	return &td_dup
}

func (td* TrainData) Subset(pos uint32, length uint32) (*TrainData) {
	var td_sub TrainData
	td_sub.object = C.fann_subset_train_data(td.object, C.uint(pos), C.uint(length))
	return &td_sub
}

func (td *TrainData) GetNumInput() (uint32) {
	return uint32(C.fann_num_input_train_data(td.object))
}

func (td *TrainData) GetNumOutput() (uint32) {
	return uint32(C.fann_num_output_train_data(td.object))
}

func (td *TrainData) SaveTrain(filename string) () {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_save_train(td.object, cfn)
}

func (td *TrainData) SaveTrainToFixed(filename string, decimal_point uint32) ( ) {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))

	C.fann_save_train_to_fixed(td.object, cfn, C.uint(decimal_point))
}

