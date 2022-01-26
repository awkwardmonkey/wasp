// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package helloworld

import "github.com/iotaledger/wasp/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableGetHelloWorldResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetHelloWorldResults) HelloWorld() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ResultHelloWorld))
}

type MutableGetHelloWorldResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetHelloWorldResults) HelloWorld() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ResultHelloWorld))
}
