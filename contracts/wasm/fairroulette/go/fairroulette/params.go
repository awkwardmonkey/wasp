// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairroulette

import "github.com/iotaledger/wasp/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutablePlaceBetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutablePlaceBetParams) Number() wasmtypes.ScImmutableUint16 {
	return wasmtypes.NewScImmutableUint16(s.proxy.Root(ParamNumber))
}

type MutablePlaceBetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutablePlaceBetParams) Number() wasmtypes.ScMutableUint16 {
	return wasmtypes.NewScMutableUint16(s.proxy.Root(ParamNumber))
}

type ImmutablePlayPeriodParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutablePlayPeriodParams) PlayPeriod() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamPlayPeriod))
}

type MutablePlayPeriodParams struct {
	proxy wasmtypes.Proxy
}

func (s MutablePlayPeriodParams) PlayPeriod() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamPlayPeriod))
}
