package sandbox

import (
	"fmt"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/sctransaction"
	"github.com/iotaledger/wasp/packages/sctransaction/txbuilder"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm"
	"github.com/iotaledger/wasp/packages/vm/vmtypes"
	"github.com/iotaledger/wasp/plugins/publisher"
)

type sandbox struct {
	*vm.VMContext
	saveTxBuilder     *txbuilder.Builder // for rollback
	requestWrapper    *requestWrapper
	stateWrapper      *stateWrapper
	contractCallStack []uint16
}

func New(vctx *vm.VMContext) vmtypes.Sandbox {
	return &sandbox{
		VMContext:      vctx,
		saveTxBuilder:  vctx.TxBuilder.Clone(),
		requestWrapper: &requestWrapper{&vctx.RequestRef},
		stateWrapper: &stateWrapper{
			contractID:   vctx.ContractID,
			virtualState: vctx.VirtualState.Variables().Codec(),
			stateUpdate:  vctx.StateUpdate,
		},
		contractCallStack: make([]uint16, 0),
	}
}

// Sandbox interface

func (vctx *sandbox) Panic(v interface{}) {
	panic(v)
}

func (vctx *sandbox) Rollback() {
	vctx.TxBuilder = vctx.saveTxBuilder
	vctx.StateUpdate.Clear()
}

func (vctx *sandbox) GetContractID() coretypes.ContractID {
	return vctx.ContractID
}

func (vctx *sandbox) GetOwnerAddress() *address.Address {
	return &vctx.OwnerAddress
}

func (vctx *sandbox) GetTimestamp() int64 {
	return vctx.Timestamp
}

func (vctx *sandbox) GetEntropy() hashing.HashValue {
	return vctx.VMContext.Entropy
}

func (vctx *sandbox) GetWaspLog() *logger.Logger {
	return vctx.Log
}

func (vctx *sandbox) DumpAccount() string {
	return vctx.TxBuilder.Dump()
}

// request arguments

func (vctx *sandbox) AccessRequest() vmtypes.RequestAccess {
	return vctx.requestWrapper
}

func (vctx *sandbox) AccessState() codec.MutableMustCodec {
	return vctx.stateWrapper.MustCodec()
}

func (vctx *sandbox) AccessSCAccount() vmtypes.AccountAccess {
	return vctx
}

func (vctx *sandbox) SendRequest(par vmtypes.NewRequestParams) bool {
	if par.IncludeReward > 0 {
		availableIotas := vctx.TxBuilder.GetInputBalance(balance.ColorIOTA)
		if par.IncludeReward+1 > availableIotas {
			return false
		}
		err := vctx.TxBuilder.MoveTokensToAddress((address.Address)(par.TargetContractID.ChainID()), balance.ColorIOTA, par.IncludeReward)
		if err != nil {
			return false
		}
	}
	reqBlock := sctransaction.NewRequestBlock(vctx.mustCurrentContractIndex(), par.TargetContractID, par.EntryPoint)
	reqBlock.WithTimelock(par.Timelock)
	reqBlock.SetArgs(par.Params)

	if err := vctx.TxBuilder.AddRequestBlock(reqBlock); err != nil {
		return false
	}
	return true
}

func (vctx *sandbox) SendRequestToSelf(reqCode coretypes.EntryPointCode, args dict.Dict) bool {
	return vctx.SendRequest(vmtypes.NewRequestParams{
		TargetContractID: vctx.ContractID,
		EntryPoint:       reqCode,
		Params:           args,
		IncludeReward:    0,
	})
}

func (vctx *sandbox) SendRequestToSelfWithDelay(entryPoint coretypes.EntryPointCode, args dict.Dict, delaySec uint32) bool {
	timelock := util.NanoSecToUnixSec(vctx.Timestamp) + delaySec

	return vctx.SendRequest(vmtypes.NewRequestParams{
		TargetContractID: vctx.ContractID,
		EntryPoint:       entryPoint,
		Params:           args,
		Timelock:         timelock,
		IncludeReward:    0,
	})
}

func (vctx *sandbox) Publish(msg string) {
	vctx.Log.Infof("VMMSG: %s '%s'", vctx.ProgramHash.String(), msg)
	publisher.Publish("vmmsg", vctx.ProgramHash.String(), msg)
}

func (vctx *sandbox) Publishf(format string, args ...interface{}) {
	vctx.Log.Infof("VMMSG: "+format, args...)
	publisher.Publish("vmmsg", vctx.ProgramHash.String(), fmt.Sprintf(format, args...))
}
