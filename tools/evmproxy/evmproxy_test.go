package main

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ethlog "github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/iotaledger/wasp/packages/evm"
	"github.com/iotaledger/wasp/packages/evm/evmtest"
	"github.com/iotaledger/wasp/tools/evmproxy/service"
	"github.com/stretchr/testify/require"
)

type env struct {
	t      *testing.T
	server *rpc.Server
	client *ethclient.Client
}

func newEnv(t *testing.T) *env {
	ethlog.Root().SetHandler(ethlog.FuncHandler(func(r *ethlog.Record) error {
		if r.Lvl <= ethlog.LvlWarn {
			t.Logf("[%s] %s", r.Lvl.AlignedString(), r.Msg)
		}
		return nil
	}))

	solo := service.NewSoloBackend(core.GenesisAlloc{
		faucetAddress: {Balance: faucetSupply},
	})
	soloEVMChain := service.NewEVMChain(solo)

	signer, _ := solo.Env.NewKeyPairWithFunds()

	rpcsrv := NewRPCServer(soloEVMChain, signer)
	t.Cleanup(rpcsrv.Stop)

	client := ethclient.NewClient(rpc.DialInProc(rpcsrv))
	t.Cleanup(client.Close)

	return &env{t, rpcsrv, client}
}

func generateKey(t *testing.T) (*ecdsa.PrivateKey, common.Address) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	return key, addr
}

var requestFundsAmount = big.NewInt(1e18) // 1 ETH

func (e *env) requestFunds(target common.Address) *types.Transaction {
	nonce, err := e.client.NonceAt(context.Background(), faucetAddress, nil)
	require.NoError(e.t, err)
	tx, err := types.SignTx(
		types.NewTransaction(nonce, target, requestFundsAmount, evm.TxGas, evm.GasPrice, nil),
		evm.Signer(),
		faucetKey,
	)
	require.NoError(e.t, err)
	err = e.client.SendTransaction(context.Background(), tx)
	require.NoError(e.t, err)
	return tx
}

func (e *env) deployEVMContract(creator *ecdsa.PrivateKey, contractABI abi.ABI, contractBytecode []byte, args ...interface{}) (*types.Transaction, common.Address) {
	creatorAddress := crypto.PubkeyToAddress(creator.PublicKey)

	nonce := e.nonceAt(creatorAddress)

	constructorArguments, err := contractABI.Pack("", args...)
	require.NoError(e.t, err)

	data := append(contractBytecode, constructorArguments...)

	value := big.NewInt(0)

	gasLimit := e.estimateGas(ethereum.CallMsg{
		From:     creatorAddress,
		To:       nil, // contract creation
		Gas:      evm.MaxGasLimit,
		GasPrice: evm.GasPrice,
		Value:    value,
		Data:     data,
	})

	tx, err := types.SignTx(
		types.NewContractCreation(nonce, value, gasLimit, evm.GasPrice, data),
		evm.Signer(),
		creator,
	)
	require.NoError(e.t, err)

	err = e.client.SendTransaction(context.Background(), tx)
	require.NoError(e.t, err)

	return tx, crypto.CreateAddress(creatorAddress, nonce)
}

func (e *env) estimateGas(msg ethereum.CallMsg) uint64 {
	gas, err := e.client.EstimateGas(context.Background(), msg)
	require.NoError(e.t, err)
	return gas
}

func (e *env) nonceAt(address common.Address) uint64 {
	nonce, err := e.client.NonceAt(context.Background(), address, nil)
	require.NoError(e.t, err)
	return nonce
}

func (e *env) blockNumber() uint64 {
	blockNumber, err := e.client.BlockNumber(context.Background())
	require.NoError(e.t, err)
	return blockNumber
}

func (e *env) blockByNumber(number *big.Int) *types.Block {
	block, err := e.client.BlockByNumber(context.Background(), number)
	require.NoError(e.t, err)
	return block
}

func (e *env) blockByHash(hash common.Hash) *types.Block {
	block, err := e.client.BlockByHash(context.Background(), hash)
	if err == ethereum.NotFound {
		return nil
	}
	require.NoError(e.t, err)
	return block
}

func (e *env) blockTransactionCountByHash(hash common.Hash) uint {
	n, err := e.client.TransactionCount(context.Background(), hash)
	require.NoError(e.t, err)
	return n
}

func (e *env) blockTransactionCountByNumber() uint {
	// the client only supports calling this method with "pending"
	n, err := e.client.PendingTransactionCount(context.Background())
	require.NoError(e.t, err)
	return n
}

func (e *env) balance(address common.Address) *big.Int {
	bal, err := e.client.BalanceAt(context.Background(), address, nil)
	require.NoError(e.t, err)
	return bal
}

func (e *env) code(address common.Address) []byte {
	code, err := e.client.CodeAt(context.Background(), address, nil)
	require.NoError(e.t, err)
	return code
}

func (e *env) storage(address common.Address, key common.Hash) []byte {
	data, err := e.client.StorageAt(context.Background(), address, key, nil)
	require.NoError(e.t, err)
	return data
}

func (e *env) txReceipt(hash common.Hash) *types.Receipt {
	r, err := e.client.TransactionReceipt(context.Background(), hash)
	require.NoError(e.t, err)
	return r
}

func TestRPCGetBalance(t *testing.T) {
	env := newEnv(t)
	_, receiverAddress := generateKey(t)
	require.Zero(t, big.NewInt(0).Cmp(env.balance(receiverAddress)))
	env.requestFunds(receiverAddress)
	require.Zero(t, big.NewInt(1e18).Cmp(env.balance(receiverAddress)))
}

func TestRPCGetCode(t *testing.T) {
	env := newEnv(t)
	creator, creatorAddress := generateKey(t)

	// account address
	{
		env.requestFunds(creatorAddress)
		require.Empty(t, env.code(creatorAddress))
	}
	// contract address
	{
		contractABI, err := abi.JSON(strings.NewReader(evmtest.StorageContractABI))
		require.NoError(t, err)
		_, contractAddress := env.deployEVMContract(creator, contractABI, evmtest.StorageContractBytecode, uint32(42))
		require.NotEmpty(t, env.code(contractAddress))
	}
}

func TestRPCGetStorage(t *testing.T) {
	env := newEnv(t)
	creator, creatorAddress := generateKey(t)

	env.requestFunds(creatorAddress)

	contractABI, err := abi.JSON(strings.NewReader(evmtest.StorageContractABI))
	require.NoError(t, err)
	_, contractAddress := env.deployEVMContract(creator, contractABI, evmtest.StorageContractBytecode, uint32(42))

	// first static variable in contract (uint32 n) has slot 0. See:
	// https://docs.soliditylang.org/en/v0.6.6/miscellaneous.html#layout-of-state-variables-in-storage
	slot := common.Hash{}
	ret := env.storage(contractAddress, slot)

	var v uint32
	err = contractABI.UnpackIntoInterface(&v, "retrieve", ret)
	require.NoError(t, err)
	require.Equal(t, uint32(42), v)
}

func TestRPCBlockNumber(t *testing.T) {
	env := newEnv(t)
	_, receiverAddress := generateKey(t)
	require.EqualValues(t, 0, env.blockNumber())
	env.requestFunds(receiverAddress)
	require.EqualValues(t, 1, env.blockNumber())
}

func TestRPCGetTransactionCount(t *testing.T) {
	env := newEnv(t)
	_, receiverAddress := generateKey(t)
	require.EqualValues(t, 0, env.nonceAt(faucetAddress))
	env.requestFunds(receiverAddress)
	require.EqualValues(t, 1, env.nonceAt(faucetAddress))
}

func TestRPCGetBlockByNumber(t *testing.T) {
	env := newEnv(t)
	_, receiverAddress := generateKey(t)
	require.EqualValues(t, 0, env.blockByNumber(big.NewInt(0)).Number().Uint64())
	env.requestFunds(receiverAddress)
	require.EqualValues(t, 1, env.blockByNumber(big.NewInt(1)).Number().Uint64())
}

func TestRPCGetBlockByHash(t *testing.T) {
	env := newEnv(t)
	_, receiverAddress := generateKey(t)
	require.Nil(t, env.blockByHash(common.Hash{}))
	require.EqualValues(t, 0, env.blockByHash(env.blockByNumber(big.NewInt(0)).Hash()).Number().Uint64())
	env.requestFunds(receiverAddress)
	require.EqualValues(t, 1, env.blockByHash(env.blockByNumber(big.NewInt(1)).Hash()).Number().Uint64())
}

func TestRPCGetTransactionCountByHash(t *testing.T) {
	env := newEnv(t)
	_, receiverAddress := generateKey(t)
	env.requestFunds(receiverAddress)
	block1 := env.blockByNumber(big.NewInt(1))
	require.Positive(t, len(block1.Transactions()))
	require.EqualValues(t, len(block1.Transactions()), env.blockTransactionCountByHash(block1.Hash()))
	require.EqualValues(t, 0, env.blockTransactionCountByHash(common.Hash{}))
}

func TestRPCGetTransactionCountByNumber(t *testing.T) {
	env := newEnv(t)
	_, receiverAddress := generateKey(t)
	env.requestFunds(receiverAddress)
	block1 := env.blockByNumber(big.NewInt(1))
	require.Positive(t, len(block1.Transactions()))
	require.EqualValues(t, len(block1.Transactions()), env.blockTransactionCountByNumber())
}

func TestRPCGetTxReceipt(t *testing.T) {
	env := newEnv(t)
	creator, creatorAddr := generateKey(t)

	// regular transaction
	{
		tx := env.requestFunds(creatorAddr)
		receipt := env.txReceipt(tx.Hash())

		require.EqualValues(t, types.LegacyTxType, receipt.Type)
		require.EqualValues(t, types.ReceiptStatusSuccessful, receipt.Status)
		require.NotZero(t, receipt.CumulativeGasUsed)
		require.EqualValues(t, types.Bloom{}, receipt.Bloom)
		require.EqualValues(t, 0, len(receipt.Logs))

		require.EqualValues(t, tx.Hash(), receipt.TxHash)
		require.EqualValues(t, common.Address{}, receipt.ContractAddress)
		require.NotZero(t, receipt.GasUsed)

		require.EqualValues(t, big.NewInt(1), receipt.BlockNumber)
		require.EqualValues(t, env.blockByNumber(big.NewInt(1)).Hash(), receipt.BlockHash)
		require.EqualValues(t, 0, receipt.TransactionIndex)
	}

	// contract creation
	{
		contractABI, err := abi.JSON(strings.NewReader(evmtest.StorageContractABI))
		require.NoError(t, err)
		tx, contractAddress := env.deployEVMContract(creator, contractABI, evmtest.StorageContractBytecode, uint32(42))
		receipt := env.txReceipt(tx.Hash())

		require.EqualValues(t, types.LegacyTxType, receipt.Type)
		require.EqualValues(t, types.ReceiptStatusSuccessful, receipt.Status)
		require.NotZero(t, receipt.CumulativeGasUsed)
		require.EqualValues(t, types.Bloom{}, receipt.Bloom)
		require.EqualValues(t, 0, len(receipt.Logs))

		require.EqualValues(t, tx.Hash(), receipt.TxHash)
		require.EqualValues(t, contractAddress, receipt.ContractAddress)
		require.NotZero(t, receipt.GasUsed)

		require.EqualValues(t, big.NewInt(2), receipt.BlockNumber)
		require.EqualValues(t, env.blockByNumber(big.NewInt(2)).Hash(), receipt.BlockHash)
		require.EqualValues(t, 0, receipt.TransactionIndex)
	}
}

func TestRPCCall(t *testing.T) {
	env := newEnv(t)
	creator, creatorAddress := generateKey(t)
	contractABI, err := abi.JSON(strings.NewReader(evmtest.StorageContractABI))
	require.NoError(t, err)
	_, contractAddress := env.deployEVMContract(creator, contractABI, evmtest.StorageContractBytecode, uint32(42))

	callArguments, err := contractABI.Pack("retrieve")
	require.NoError(t, err)

	ret, err := env.client.CallContract(context.Background(), ethereum.CallMsg{
		From: creatorAddress,
		To:   &contractAddress,
		Data: callArguments,
	}, nil)
	require.NoError(t, err)

	var v uint32
	err = contractABI.UnpackIntoInterface(&v, "retrieve", ret)
	require.NoError(t, err)
	require.Equal(t, uint32(42), v)
}
