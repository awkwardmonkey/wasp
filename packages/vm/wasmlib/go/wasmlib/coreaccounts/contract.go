// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreaccounts

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type DepositCall struct {
	Func   *wasmlib.ScFunc
	Params MutableDepositParams
}

type HarvestCall struct {
	Func   *wasmlib.ScFunc
	Params MutableHarvestParams
}

type WithdrawCall struct {
	Func *wasmlib.ScFunc
}

type AccountsCall struct {
	Func    *wasmlib.ScView
	Results ImmutableAccountsResults
}

type BalanceCall struct {
	Func    *wasmlib.ScView
	Params  MutableBalanceParams
	Results ImmutableBalanceResults
}

type GetAccountNonceCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetAccountNonceParams
	Results ImmutableGetAccountNonceResults
}

type TotalAssetsCall struct {
	Func    *wasmlib.ScView
	Results ImmutableTotalAssetsResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) Deposit(ctx wasmlib.ScFuncCallContext) *DepositCall {
	f := &DepositCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncDeposit)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) Harvest(ctx wasmlib.ScFuncCallContext) *HarvestCall {
	f := &HarvestCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncHarvest)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) Withdraw(ctx wasmlib.ScFuncCallContext) *WithdrawCall {
	return &WithdrawCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncWithdraw)}
}

func (sc Funcs) Accounts(ctx wasmlib.ScViewCallContext) *AccountsCall {
	f := &AccountsCall{Func: wasmlib.NewScView(ctx, HScName, HViewAccounts)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) Balance(ctx wasmlib.ScViewCallContext) *BalanceCall {
	f := &BalanceCall{Func: wasmlib.NewScView(ctx, HScName, HViewBalance)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) GetAccountNonce(ctx wasmlib.ScViewCallContext) *GetAccountNonceCall {
	f := &GetAccountNonceCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetAccountNonce)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) TotalAssets(ctx wasmlib.ScViewCallContext) *TotalAssetsCall {
	f := &TotalAssetsCall{Func: wasmlib.NewScView(ctx, HScName, HViewTotalAssets)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncDeposit, wasmlib.FuncError)
	exports.AddFunc(FuncHarvest, wasmlib.FuncError)
	exports.AddFunc(FuncWithdraw, wasmlib.FuncError)
	exports.AddView(ViewAccounts, wasmlib.ViewError)
	exports.AddView(ViewBalance, wasmlib.ViewError)
	exports.AddView(ViewGetAccountNonce, wasmlib.ViewError)
	exports.AddView(ViewTotalAssets, wasmlib.ViewError)
}