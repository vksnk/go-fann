package fann
/*
#include <fann.h>
*/
import "C"

//list all activation functions
type ActivationFunc C.enum_fann_activationfunc_enum


var LINEAR ActivationFunc = C.FANN_LINEAR
var THRESHOLD ActivationFunc = C.FANN_THRESHOLD
var THRESHOLD_SYMMETRIC ActivationFunc = C.FANN_THRESHOLD_SYMMETRIC
var SIGMOID ActivationFunc = C.FANN_SIGMOID
var SIGMOID_STEPWISE ActivationFunc = C.FANN_SIGMOID_STEPWISE
var SIGMOID_SYMMETRIC ActivationFunc = C.FANN_SIGMOID_SYMMETRIC
var SIGMOID_SYMMETRIC_STEPWISE ActivationFunc = C.FANN_SIGMOID_SYMMETRIC_STEPWISE
var GAUSSIAN ActivationFunc = C.FANN_GAUSSIAN
var GAUSSIAN_SYMMETRIC ActivationFunc = C.FANN_GAUSSIAN_SYMMETRIC
var GAUSSIAN_STEPWISE ActivationFunc = C.FANN_GAUSSIAN_STEPWISE
var ELLIOT ActivationFunc = C.FANN_ELLIOT
var ELLIOT_SYMMETRIC ActivationFunc = C.FANN_ELLIOT_SYMMETRIC
var LINEAR_PIECE ActivationFunc = C.FANN_LINEAR_PIECE
var LINEAR_PIECE_SYMMETRIC ActivationFunc = C.FANN_LINEAR_PIECE_SYMMETRIC
var SIN_SYMMETRIC ActivationFunc = C.FANN_SIN_SYMMETRIC
var COS_SYMMETRIC ActivationFunc = C.FANN_COS_SYMMETRIC
var SIN ActivationFunc = C.FANN_SIN
var COS ActivationFunc = C.FANN_COS

//net types
type Nettype C.enum_fann_nettype_enum

var FANN_NETTYPE_LAYER Nettype = C.FANN_NETTYPE_LAYER
var FANN_NETTYPE_SHORTCUT Nettype = C.FANN_NETTYPE_SHORTCUT


