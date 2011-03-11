package fann

/*
#include <fann.h>

typedef unsigned int* uint_ptr;

const unsigned int* go2uintArray(unsigned int* arr, int n) {
	return malloc(sizeof(unsigned int) * n);
}

*/
import "C"
import "unsafe"

type FannType C.fann_type

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

func CreateFromFile(filename string) (*Ann) {
	var ann Ann
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	ann.object = C.fann_create_from_file(cfn)
	return &ann
}

func (ann *Ann) TrainOnFile(filename string, maxEpoches uint32, epochBetweenReports uint32, desiredError float32) {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_train_on_file(ann.object, cfn, C.uint(maxEpoches), C.uint(epochBetweenReports), C.float(desiredError));
}

func (ann *Ann) Run(input []FannType) ([]FannType) {
	c_out := C.fann_run(ann.object, (*C.fann_type)(&input[0]))
	out := make([]FannType, 1)
	out[0] = FannType(*c_out)
	return out
}

func (ann *Ann) Save(filename string) {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_save(ann.object, cfn)
}

func (ann *Ann) Destroy() {
	C.fann_destroy(ann.object)
}

func (*Ann) Foo() {

}
/*
func CreateSparse(concentration float32, num_layers uint, layers []int) {
}
*/
