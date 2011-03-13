package fann
/*
#include <fann.h>
*/
import "C"

type Error struct {
	object *C.struct_fann_error
}
/*
//FANN_EXTERNAL enum fann_errno_enum FANN_API fann_get_errno(struct fann_error *errdat);
func GetErrno(errdat []) ( ) {
	C.fann_get_errno()
}
*/
func (err *Error) ResetErrno() ( ) {
	C.fann_reset_errno(err.object)
}

//FANN_EXTERNAL void FANN_API fann_reset_errstr(struct fann_error *errdat);
func (err *Error) ResetErrstr() ( ) {
	C.fann_reset_errstr(err.object)
}
/*
//FANN_EXTERNAL char *FANN_API fann_get_errstr(struct fann_error *errdat);
func GetErrstr() () {
	C.fann_get_errstr()
}
*/
func (err *Error) PrintError() ( ) {
	C.fann_print_error(err.object)
}

