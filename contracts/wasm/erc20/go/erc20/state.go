// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasp/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type MapAgentIDToImmutableAllowancesForAgent struct {
	proxy wasmtypes.Proxy
}

func (m MapAgentIDToImmutableAllowancesForAgent) GetAllowancesForAgent(key wasmtypes.ScAgentID) ImmutableAllowancesForAgent {
	return ImmutableAllowancesForAgent{proxy: m.proxy.Key(wasmtypes.BytesFromAgentID(key))}
}

type ImmutableErc20State struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableErc20State) AllAllowances() MapAgentIDToImmutableAllowancesForAgent {
	return MapAgentIDToImmutableAllowancesForAgent{proxy: s.proxy.Root(StateAllAllowances)}
}

func (s ImmutableErc20State) Balances() MapAgentIDToImmutableUint64 {
	return MapAgentIDToImmutableUint64{proxy: s.proxy.Root(StateBalances)}
}

func (s ImmutableErc20State) Supply() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(StateSupply))
}

type MapAgentIDToMutableAllowancesForAgent struct {
	proxy wasmtypes.Proxy
}

func (m MapAgentIDToMutableAllowancesForAgent) Clear() {
	m.proxy.ClearMap()
}

func (m MapAgentIDToMutableAllowancesForAgent) GetAllowancesForAgent(key wasmtypes.ScAgentID) MutableAllowancesForAgent {
	return MutableAllowancesForAgent{proxy: m.proxy.Key(wasmtypes.BytesFromAgentID(key))}
}

type MutableErc20State struct {
	proxy wasmtypes.Proxy
}

func (s MutableErc20State) AsImmutable() ImmutableErc20State {
	return ImmutableErc20State(s)
}

func (s MutableErc20State) AllAllowances() MapAgentIDToMutableAllowancesForAgent {
	return MapAgentIDToMutableAllowancesForAgent{proxy: s.proxy.Root(StateAllAllowances)}
}

func (s MutableErc20State) Balances() MapAgentIDToMutableUint64 {
	return MapAgentIDToMutableUint64{proxy: s.proxy.Root(StateBalances)}
}

func (s MutableErc20State) Supply() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(StateSupply))
}
