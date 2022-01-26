// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "testwasmlib"
	ScDescription = "Exercise several aspects of WasmLib"
	HScName       = wasmtypes.ScHname(0x89703a45)
)

const (
	ParamAddress     = "address"
	ParamAgentID     = "agentID"
	ParamBlockIndex  = "blockIndex"
	ParamBool        = "bool"
	ParamBytes       = "bytes"
	ParamChainID     = "chainID"
	ParamColor       = "color"
	ParamHash        = "hash"
	ParamHname       = "hname"
	ParamIndex       = "index"
	ParamInt16       = "int16"
	ParamInt32       = "int32"
	ParamInt64       = "int64"
	ParamInt8        = "int8"
	ParamKey         = "key"
	ParamName        = "name"
	ParamParam       = "this"
	ParamRecordIndex = "recordIndex"
	ParamRequestID   = "requestID"
	ParamString      = "string"
	ParamUint16      = "uint16"
	ParamUint32      = "uint32"
	ParamUint64      = "uint64"
	ParamUint8       = "uint8"
	ParamValue       = "value"
)

const (
	ResultCount  = "count"
	ResultIotas  = "iotas"
	ResultLength = "length"
	ResultRandom = "random"
	ResultRecord = "record"
	ResultValue  = "value"
)

const (
	StateArrays = "arrays"
	StateMaps   = "maps"
	StateRandom = "random"
)

const (
	FuncArrayAppend  = "arrayAppend"
	FuncArrayClear   = "arrayClear"
	FuncArraySet     = "arraySet"
	FuncMapClear     = "mapClear"
	FuncMapSet       = "mapSet"
	FuncParamTypes   = "paramTypes"
	FuncRandom       = "random"
	FuncTriggerEvent = "triggerEvent"
	ViewArrayLength  = "arrayLength"
	ViewArrayValue   = "arrayValue"
	ViewBlockRecord  = "blockRecord"
	ViewBlockRecords = "blockRecords"
	ViewGetRandom    = "getRandom"
	ViewIotaBalance  = "iotaBalance"
	ViewMapValue     = "mapValue"
)

const (
	HFuncArrayAppend  = wasmtypes.ScHname(0x612f835f)
	HFuncArrayClear   = wasmtypes.ScHname(0x88021821)
	HFuncArraySet     = wasmtypes.ScHname(0x2c4150b3)
	HFuncMapClear     = wasmtypes.ScHname(0x027f215a)
	HFuncMapSet       = wasmtypes.ScHname(0xf2260404)
	HFuncParamTypes   = wasmtypes.ScHname(0x6921c4cd)
	HFuncRandom       = wasmtypes.ScHname(0xe86c97ca)
	HFuncTriggerEvent = wasmtypes.ScHname(0xd5438ac6)
	HViewArrayLength  = wasmtypes.ScHname(0x3a831021)
	HViewArrayValue   = wasmtypes.ScHname(0x662dbd81)
	HViewBlockRecord  = wasmtypes.ScHname(0xad13b2f8)
	HViewBlockRecords = wasmtypes.ScHname(0x16e249ea)
	HViewGetRandom    = wasmtypes.ScHname(0x46263045)
	HViewIotaBalance  = wasmtypes.ScHname(0x9d3920bd)
	HViewMapValue     = wasmtypes.ScHname(0x23149bef)
)
