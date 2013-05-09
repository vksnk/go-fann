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
func CreateStandard(numLayers uint, layers []uint32) (*Ann) {
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

func (ann *Ann) TrainEpoch(td *TrainData) (float32) {
	return float32(C.fann_train_epoch(ann.object, td.object))
}

func (ann *Ann) TestData(td *TrainData) (float32) {
	return float32(C.fann_test_data(ann.object, td.object))
}

func (ann *Ann) TrainOnData(td *TrainData, max_epochs uint32, epochs_between_reports uint32, desired_error float32) () {
	C.fann_train_on_data(ann.object, td.object, C.uint(max_epochs), C.uint(epochs_between_reports), C.float(desired_error))
}

func (ann *Ann) TrainOnFile(filename string, maxEpoches uint32, epochBetweenReports uint32, desiredError float32) {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_train_on_file(ann.object, cfn, C.uint(maxEpoches), C.uint(epochBetweenReports), C.float(desiredError));
}

func (ann *Ann) Test(input []FannType,  desired_output []FannType) ([]FannType) {
	c_out := C.fann_test(ann.object, (*C.fann_type)(&input[0]), (*C.fann_type)(&desired_output[0]))
	outputNum := ann.GetNumOutput()
	out := make([]FannType, outputNum)
	C.cpFannTypeArray(c_out, (*C.fann_type)(&out[0]), C.uint(outputNum))
	return out
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
//TODO: finish it
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

func (ann *Ann) GetQuickpropDecay() (float32) {
	return float32(C.fann_get_quickprop_decay(ann.object))
}

func (ann *Ann) SetQuickpropDecay(quickprop_decay float32) () {
	C.fann_set_quickprop_decay(ann.object, C.float(quickprop_decay))
}

func (ann *Ann) GetQuickpropMu() (float32) {
	return float32(C.fann_get_quickprop_mu(ann.object))
}

func (ann *Ann) SetQuickpropMu(quickprop_mu float32) ( ) {
	C.fann_set_quickprop_mu(ann.object, C.float(quickprop_mu))
}

func (ann *Ann) GetRpropIncreaseFactor() (float32) {
	return float32(C.fann_get_rprop_increase_factor(ann.object))
}

func (ann *Ann) SetRpropIncreaseFactor(rprop_increase_factor float32) ( ) {
	C.fann_set_rprop_increase_factor(ann.object, C.float(rprop_increase_factor))
}

func (ann *Ann) GetRpropDecreaseFactor() (float32) {
	return float32(C.fann_get_rprop_decrease_factor(ann.object))
}

func (ann *Ann) SetRpropDecreaseFactor(rprop_decrease_factor float32) ( ) {
	C.fann_set_rprop_decrease_factor(ann.object, C.float(rprop_decrease_factor))
}

func (ann *Ann) GetRpropDeltaMin() (float32) {
	return float32(C.fann_get_rprop_delta_min(ann.object))
}

func (ann *Ann) SetRpropDeltaMin(rprop_delta_min float32) ( ) {
	C.fann_set_rprop_delta_min(ann.object, C.float(rprop_delta_min))
}

func (ann *Ann) GetRpropDeltaMax() (float32) {
	return float32(C.fann_get_rprop_delta_max(ann.object))
}

func (ann *Ann) SetRpropDeltaMax(rprop_delta_max float32) ( ) {
	C.fann_set_rprop_delta_max(ann.object, C.float(rprop_delta_max))
}

func (ann *Ann) GetRpropDeltaZero() (float32) {
	return float32(C.fann_get_rprop_delta_zero(ann.object))
}

func (ann *Ann) SetRpropDeltaZero(rprop_delta_max float32) ( ) {
	C.fann_set_rprop_delta_zero(ann.object, C.float(rprop_delta_max))
}

func (ann *Ann) GetBitFailLimit() (FannType) {
	return FannType(C.fann_get_bit_fail_limit(ann.object))
}

func (ann *Ann) SetBitFailLimit(bit_fail_limit FannType) ( ) {
	C.fann_set_bit_fail_limit(ann.object, C.fann_type(bit_fail_limit))
}

func (ann *Ann) GetLearningRate() (float32) {
	return float32(C.fann_get_learning_rate(ann.object))
}

func (ann *Ann) SetLearningRate(learning_rate float32) ( ) {
	C.fann_set_learning_rate(ann.object, C.float(learning_rate))
}

func (ann *Ann) GetLearningMomentum() (float32 ) {
	return float32(C.fann_get_learning_momentum(ann.object))
}

func (ann *Ann) SetLearningMomentum(learning_momentum float32) ( ) {
	C.fann_set_learning_momentum(ann.object, C.float(learning_momentum))
}

func (ann *Ann) ScaleInput(input_vector []FannType) ( ) {
	C.fann_scale_input(ann.object, (*C.fann_type)(&input_vector[0]))
}

func (ann *Ann) ScaleOutput(output_vector []FannType) ( ) {
	C.fann_scale_output(ann.object, (*C.fann_type)(&output_vector[0]))
}

func (ann *Ann) DescaleInput(input_vector []FannType) ( ) {
	C.fann_descale_input(ann.object, (*C.fann_type)(&input_vector[0]))
}

func (ann *Ann) DescaleOutput(output_vector []FannType) ( ) {
	C.fann_descale_output(ann.object, (*C.fann_type)(&output_vector[0]))
}

func (ann *Ann) GetTrainingAlgorithm() (TrainingAlgorithm) {
	return TrainingAlgorithm(C.fann_get_training_algorithm(ann.object))
}

func (ann *Ann) SetTrainingAlgorithm(training_algorithm TrainingAlgorithm) () {
	C.fann_set_training_algorithm(ann.object, C.enum_fann_train_enum(training_algorithm))
}

func (ann *Ann) GetTrainStopFunction() (StopFunction) {
	return StopFunction(C.fann_get_train_stop_function(ann.object))
}

func (ann *Ann) SetTrainStopFunction(train_stop_function StopFunction) () {
	C.fann_set_train_stop_function(ann.object, C.enum_fann_stopfunc_enum(train_stop_function))
}

func (ann *Ann) GetTrainErrorFunction() (TrainErrorFunction) {
	return TrainErrorFunction(C.fann_get_train_error_function(ann.object))
}

func (ann *Ann) SetTrainErrorFunction(train_error_function TrainErrorFunction) () {
	C.fann_set_train_error_function(ann.object, C.enum_fann_errorfunc_enum(train_error_function))
}

func (ann *Ann) SetInputScalingParams(td *TrainData, new_input_min float32, new_input_max float32) (int) {
	return int(C.fann_set_input_scaling_params(ann.object, td.object, C.float(new_input_min), C.float(new_input_max)))
}

func (ann *Ann) SetOutputScalingParams(td *TrainData, new_output_min float32, new_output_max float32) (int) {
	return int(C.fann_set_output_scaling_params(ann.object, td.object, C.float(new_output_min), C.float(new_output_max)))
}

func (ann *Ann) SetScalingParams(td *TrainData, new_input_min float32, new_input_max float32, new_output_min float32, new_output_max float32) (int) {
	return int(C.fann_set_scaling_params(ann.object, td.object, C.float(new_input_min), C.float(new_input_max), C.float(new_output_min), C.float(new_output_max)))
}

func (ann *Ann) ClearScalingParams() (int) {
	return int(C.fann_clear_scaling_params(ann.object))
}

func (ann *Ann) GetActivationFunction(layer int, neuron int) (ActivationFunc) {
	return ActivationFunc(C.fann_get_activation_function(ann.object, C.int(layer), C.int(neuron)))
}

func (ann *Ann) SetActivationFunction(activation_function ActivationFunc, layer int, neuron int) ( ) {
	C.fann_set_activation_function(ann.object, C.enum_fann_activationfunc_enum(activation_function), C.int(layer), C.int(neuron))
}

func (ann *Ann) SetActivationFunctionLayer(activation_function ActivationFunc, layer int) () {
	C.fann_set_activation_function_layer(ann.object, C.enum_fann_activationfunc_enum(activation_function), C.int(layer))
}

func (ann *Ann) GetActivationSteepness(layer int, neuron int) (FannType) {
	return FannType(C.fann_get_activation_steepness(ann.object, C.int(layer), C.int(layer)))
}

func (ann *Ann) SetActivationSteepness(steepness FannType, layer int, neuron int) () {
	C.fann_set_activation_steepness(ann.object, C.fann_type(steepness), C.int(layer), C.int(layer))
}

func (ann *Ann) SetActivationSteepnessLayer(steepness FannType, layer int) ( ) {
	C.fann_set_activation_steepness_layer(ann.object, C.fann_type(steepness), C.int(layer))
}

func (ann *Ann) SetActivationSteepnessHidden(steepness FannType) ( ) {
	C.fann_set_activation_steepness_hidden(ann.object, C.fann_type(steepness))
}

func (ann *Ann) SetActivationSteepnessOutput(steepness FannType) () {
	C.fann_set_activation_steepness_output(ann.object, C.fann_type(steepness))
}
