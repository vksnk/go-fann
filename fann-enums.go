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

var NETTYPE_LAYER Nettype = C.FANN_NETTYPE_LAYER
var NETTYPE_SHORTCUT Nettype = C.FANN_NETTYPE_SHORTCUT

//training algs
type TrainingAlgorithm C.enum_fann_train_enum

var TRAIN_INCREMENTAL TrainingAlgorithm = C.FANN_TRAIN_INCREMENTAL
var TRAIN_BATCH TrainingAlgorithm = C.FANN_TRAIN_BATCH
var TRAIN_RPROP TrainingAlgorithm = C.FANN_TRAIN_RPROP
var TRAIN_QUICKPROP TrainingAlgorithm = C.FANN_TRAIN_QUICKPROP

//stop functions
type StopFunction C.enum_fann_stopfunc_enum

var STOPFUNC_MSE StopFunction = C.FANN_STOPFUNC_MSE
var STOPFUNC_BIT StopFunction = C.FANN_STOPFUNC_BIT

