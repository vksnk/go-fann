package fann

/*
#include <fann.h>

static void cpFuncActArray(enum fann_activationfunc_enum* src, enum fann_activationfunc_enum* dst, unsigned int n) {
	unsigned int i = 0;
	for(; i < n; i++)
		dst[i] = src[i];
}

static void cpFannTypeArray(fann_type* src, fann_type* dst, unsigned int n) {
	unsigned int i = 0;
	for(; i < n; i++)
		dst[i] = src[i];
}

*/
import "C"
import "unsafe"

func (ann *Ann) CascadeTrainOnData(td *TrainData, max_neurons uint32, neurons_between_reports uint32, desired_error float32) () {
	C.fann_cascadetrain_on_data(ann.object, td.object, C.uint(max_neurons), C.uint(neurons_between_reports), C.float(desired_error))
}

func (ann *Ann) CascadeTrainOnFile(filename string, max_neurons uint32, neurons_between_reports uint32, desired_error float32) ( ) {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_cascadetrain_on_file(ann.object, cfn, C.uint(max_neurons), C.uint(neurons_between_reports), C.float(desired_error))
}

func (ann *Ann) GetCascadeOutputChangeFraction() (float32) {
	return float32(C.fann_get_cascade_output_change_fraction(ann.object))
}

func (ann *Ann) SetCascadeOutputChangeFraction(cascade_output_change_fraction float32) ( ) {
	C.fann_set_cascade_output_change_fraction(ann.object, C.float(cascade_output_change_fraction))
}

func (ann *Ann) GetCascadeNumCandidateGroups() (uint32) {
	return uint32(C.fann_get_cascade_num_candidate_groups(ann.object))
}

func (ann *Ann) SetCascadeNumCandidateGroups(cascade_num_candidate_groups uint32) () {
	C.fann_set_cascade_num_candidate_groups(ann.object, C.uint(cascade_num_candidate_groups))
}


func (ann *Ann) GetCascadeOutputStagnationEpochs() (uint32) {
	return uint32(C.fann_get_cascade_output_stagnation_epochs(ann.object))
}

func (ann *Ann) SetCascadeOutputStagnationEpochs(cascade_output_stagnation_epochs uint32) ( ) {
	C.fann_set_cascade_output_stagnation_epochs(ann.object, C.uint(cascade_output_stagnation_epochs))
}

func (ann *Ann) GetCascadeCandidateChangeFraction() (float32) {
	return float32(C.fann_get_cascade_candidate_change_fraction(ann.object))
}

func (ann *Ann) SetCascadeCandidateChangeFraction(cascade_candidate_change_fraction float32) ( ) {
	C.fann_set_cascade_candidate_change_fraction(ann.object, C.float(cascade_candidate_change_fraction))
}

func (ann *Ann) GetCascadeCandidateStagnationEpochs() (uint32) {
	return uint32(C.fann_get_cascade_candidate_stagnation_epochs(ann.object))
}

func (ann *Ann) SetCascadeCandidateStagnationEpochs(cascade_candidate_stagnation_epochs uint32) ( ) {
	C.fann_set_cascade_candidate_stagnation_epochs(ann.object, C.uint(cascade_candidate_stagnation_epochs))
}

func (ann *Ann) GetCascadeWeightMultiplier() (FannType) {
	return FannType(C.fann_get_cascade_weight_multiplier(ann.object))
}

func (ann *Ann) SetCascadeWeightMultiplier(cascade_weight_multiplier FannType) () {
	C.fann_set_cascade_weight_multiplier(ann.object, C.fann_type(cascade_weight_multiplier))
}

func (ann *Ann) GetCascadeCandidateLimit() (FannType) {
	return FannType(C.fann_get_cascade_candidate_limit(ann.object))
}

func (ann *Ann) SetCascadeCandidateLimit(cascade_candidate_limit FannType) ( ) {
	C.fann_set_cascade_candidate_limit(ann.object, C.fann_type(cascade_candidate_limit))
}

func (ann *Ann) GetCascadeMaxOutEpochs() (uint32) {
	return uint32(C.fann_get_cascade_max_out_epochs(ann.object))
}

func (ann *Ann) SetCascadeMaxOutEpochs(cascade_max_out_epochs uint32) ( ) {
	C.fann_set_cascade_max_out_epochs(ann.object, C.uint(cascade_max_out_epochs))
}

func (ann *Ann) GetCascadeMaxCandEpochs() (uint32) {
	return uint32(C.fann_get_cascade_max_cand_epochs(ann.object))
}

func (ann *Ann) SetCascadeMaxCandEpochs(cascade_max_cand_epochs uint32) ( ) {
	C.fann_set_cascade_max_cand_epochs(ann.object, C.uint(cascade_max_cand_epochs))
}

func (ann *Ann) GetCascadeNumCandidates() (uint32) {
	return uint32(C.fann_get_cascade_num_candidates(ann.object))
}

func (ann *Ann) GetCascadeActivationFunctionsCount() (uint32) {
	return uint32(C.fann_get_cascade_activation_functions_count(ann.object))
}

func (ann *Ann) GetCascadeActivationFunctions() ([]ActivationFunc) {
	cout := C.fann_get_cascade_activation_functions(ann.object)
	n := ann.GetCascadeActivationFunctionsCount()
	out := make([]ActivationFunc, n)
	C.cpFuncActArray(cout, (*C.enum_fann_activationfunc_enum)(&out[0]), C.uint(n))
	return out
}


func (ann *Ann) SetCascadeActivationFunctions(cascade_activation_functions []ActivationFunc) ( ) {
	C.fann_set_cascade_activation_functions(ann.object, (*C.enum_fann_activationfunc_enum)(&cascade_activation_functions[0]), C.uint(len(cascade_activation_functions)))
}

func (ann *Ann) GetCascadeActivationSteepnessesCount() (uint32) {
	return uint32(C.fann_get_cascade_activation_steepnesses_count(ann.object))
}

func (ann *Ann) GetCascadeActivationSteepnesses() ([]FannType) {
	cout := C.fann_get_cascade_activation_steepnesses(ann.object)
	n := ann.GetCascadeActivationSteepnessesCount()
	out := make([]FannType, n)
	C.cpFannTypeArray(cout, (*C.fann_type)(&out[0]), C.uint(n))
	return out
}

func (ann *Ann) SetCascadeActivationSteepnesses(cascade_activation_steepnesses []FannType) ( ) {
	C.fann_set_cascade_activation_steepnesses(ann.object, (*C.fann_type)(&cascade_activation_steepnesses[0]), C.uint(len(cascade_activation_steepnesses)))
}


