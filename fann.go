package fann

/*
#include <fann.h>

typedef unsigned int* uint_ptr;

const unsigned int* go2uintArray(unsigned int* arr, int n) {
	return malloc(sizeof(unsigned int) * n);
}

static void cpFannTypeArray(fann_type* src, fann_type* dst, unsigned int n) {
	unsigned int i = 0;
	for(; i < n; i++)
		dst[i] = src[i];
}

*/
import "C"
import "unsafe"

type FannType C.fann_type

type Connection C.struct_fann_connection

type TrainData struct {
	object *C.struct_fann_train_data
}



type Ann struct {
	object *C.struct_fann
}

//Create ann functions
func CreateStandart(numLayers uint, layers []uint32) (*Ann) {
	var ann Ann
	ann.object = C.fann_create_standard_array(C.uint(numLayers), (*C.uint)(&layers[0]))
	return &ann
}

func CreateSparse(concentration float32, numLayers uint, layers []uint32) (*Ann) {
	var ann Ann
	ann.object = C.fann_create_sparse_array(C.float(concentration), C.uint(numLayers), (*C.uint)(&layers[0]))
	return &ann
}

func CreateShortcut(num_layers uint32, layers []uint32) (*Ann) {
	var ann Ann
	ann.object = C.fann_create_shortcut_array(C.uint(num_layers), (*C.uint)(&layers[0]))
	return &ann
}

func CreateFromFile(filename string) (*Ann) {
	var ann Ann
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	ann.object = C.fann_create_from_file(cfn)
	return &ann
}

//run & test functions
func (ann *Ann) Run(input []FannType) ([]FannType) {
	c_out := C.fann_run(ann.object, (*C.fann_type)(&input[0]))
	outputNum := ann.GetNumOutput()
	out := make([]FannType, outputNum)
	C.cpFannTypeArray(c_out, (*C.fann_type)(&out[0]), C.uint(outputNum))
	return out
}

func (ann *Ann) Train(input []FannType,  desired_output []FannType) ( ) {
	C.fann_train(ann.object, (*C.fann_type)(&input[0]), (*C.fann_type)(&desired_output[0]))
}

func (ann *Ann) TrainOnData(td *TrainData, max_epochs uint32, epochs_between_reports uint32, desired_error float32) () {
	C.fann_train_on_data(ann.object, td.object, C.uint(max_epochs), C.uint(epochs_between_reports), C.float(desired_error))
}

func (ann *Ann) TrainOnFile(filename string, maxEpoches uint32, epochBetweenReports uint32, desiredError float32) {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_train_on_file(ann.object, cfn, C.uint(maxEpoches), C.uint(epochBetweenReports), C.float(desiredError));
}

//TODO: finish it
func (ann *Ann) Test(input []FannType,  desired_output []FannType) ( ) {
	C.fann_test(ann.object, (*C.fann_type)(&input[0]), (*C.fann_type)(&desired_output[0]))
}

func (ann *Ann) GetMSE() (float32) {
	return float32(C.fann_get_MSE(ann.object))
}

func (ann *Ann) GetBitFail() (uint32) {
	return uint32(C.fann_get_bit_fail(ann.object))
}

func (ann *Ann) ResetMSE() () {
	C.fann_reset_MSE(ann.object)
}

func (ann *Ann) InitWeights(train_data *TrainData) () {
	C.fann_init_weights(ann.object, train_data.object)
}

func (ann *Ann) RandomizeWeights(min_weight FannType, max_weight FannType) ( ) {
	C.fann_randomize_weights(ann.object, C.fann_type(min_weight), C.fann_type(max_weight))
}
//print functions
func (ann *Ann) PrintConnections() ( ) {
	C.fann_print_connections(ann.object)
}

func (ann *Ann) PrintParameters() ( ) {
	C.fann_print_parameters(ann.object)
}

func (ann *Ann) SetActivationFunctionHidden(tp ActivationFunc) {
	C.fann_set_activation_function_hidden(ann.object, C.enum_fann_activationfunc_enum(tp))
}

func (ann *Ann) SetActivationFunctionOutput(tp ActivationFunc) {
	C.fann_set_activation_function_output(ann.object, C.enum_fann_activationfunc_enum(tp))
}
//save functions
func (ann *Ann) Save(filename string) {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_save(ann.object, cfn)
}

func (ann *Ann) SaveToFixed(configuration_file string) {
	cfn := C.CString(configuration_file)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_save_to_fixed(ann.object, cfn)
}
//destroy function
func (ann *Ann) Destroy() {
	C.fann_destroy(ann.object)
}

//getters
func (ann *Ann) GetNumInput() (uint32) {
	return uint32(C.fann_get_num_input(ann.object))
}

func (ann *Ann) GetNumOutput() (uint32) {
	return uint32(C.fann_get_num_output(ann.object))
}

func (ann *Ann) GetTotalNeurons() (uint32) {
	return uint32(C.fann_get_total_neurons(ann.object))
}

func (ann *Ann) GetTotalConnections() (uint32) {
	return uint32(C.fann_get_total_connections(ann.object))
}

func (ann *Ann) GetConnectionRate() (float32) {
	return float32(C.fann_get_connection_rate(ann.object))
}

func (ann *Ann) GetNumLayers() (uint32) {
	return uint32(C.fann_get_num_layers(ann.object))
}
/*
TODO: finish it
func (ann *Ann) GetDecimalPoint() (uint32) {
	return uint32(C.fann_get_decimal_point(ann.object))
}

func (ann *Ann) GetMultiplier() (uint32) {
	return uint32(C.fann_get_multiplier(ann.object))
}
*/

func (ann *Ann) GetNetworkType() (Nettype) {
	return Nettype(C.fann_get_network_type(ann.object))
}

func (ann *Ann) GetLayerArray() ([]uint32) {
	layers := make([]uint32, ann.GetNumLayers())
	C.fann_get_layer_array(ann.object, (*C.uint)(&layers[0]))
	return layers
}

func (ann *Ann) GetBiasArray() ([]uint32) {
	bias := make([]uint32, ann.GetNumLayers())
	C.fann_get_bias_array(ann.object, (*C.uint)(&bias[0]))
	return bias
}

/*
TODO: finish it
//FANN_EXTERNAL void FANN_API fann_get_connection_array(struct fann *ann,struct fann_connection *connections);
func (ann *Ann) GetConnectionArray(connections []Connection) ( ) {
	C.fann_get_connection_array(ann.object, )
}
*/

//setters
func (ann *Ann) SetWeightArray(connections []Connection, num_connections uint32) ( ) {
	C.fann_set_weight_array(ann.object, (*C.struct_fann_connection)(&connections[0]), C.uint(num_connections))
}

func (ann *Ann) SetWeight(from_neuron uint32, to_neuron uint32, weight FannType) ( ) {
	C.fann_set_weight(ann.object, C.uint(from_neuron), C.uint(to_neuron), C.fann_type(weight))
}

