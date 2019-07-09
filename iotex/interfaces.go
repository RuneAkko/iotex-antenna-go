package iotex

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"google.golang.org/grpc"
)

// SendActionCaller is used to perform a send action call.
type SendActionCaller interface {
	API() iotexapi.APIServiceClient
	Call(ctx context.Context, opts ...grpc.CallOption) (hash.Hash256, error)
}

// TransferCaller is used to perform a transfer call.
type TransferCaller interface {
	SendActionCaller

	SetGasPrice(*big.Int) TransferCaller
	SetGasLimit(uint64) TransferCaller
	SetPayload([]byte) TransferCaller
}

// DeployContractCaller is used to perform a deploy contract call.
type DeployContractCaller interface {
	SendActionCaller

	SetArgs(abi abi.ABI, args ...interface{}) DeployContractCaller
	SetGasPrice(*big.Int) DeployContractCaller
	SetGasLimit(uint64) DeployContractCaller
}

// GetReceiptCaller is used to perform a get receipt call.
type GetReceiptCaller interface {
	Call(ctx context.Context, opts ...grpc.CallOption) (*iotexapi.GetReceiptByActionResponse, error)
}

// GetExecutionResultCaller is used to perform a get execution result call.
type GetExecutionResultCaller interface {
	Call(ctx context.Context, opts ...grpc.CallOption) ([]byte, error)
}

// AuthedClient is an iotex client which associate with an account credentials, so it can perform write actions.
type AuthedClient interface {
	ReadOnlyClient

	Contract(contract address.Address, abi abi.ABI) Contract
	Transfer(to address.Address, value *big.Int) TransferCaller
	DeployContract(data []byte) DeployContractCaller
}

// ReadOnlyClient is an iotex client which can perform read actions.
type ReadOnlyClient interface {
	API() iotexapi.APIServiceClient

	ReadOnlyContract(contract address.Address, abi abi.ABI) ReadOnlyContract
	GetReceipt(actionHash hash.Hash256) GetReceiptCaller
	GetExecutionResult(actionHash hash.Hash256) GetExecutionResultCaller
}

// ReadContractCaller is used to perform a read contract call.
type ReadContractCaller interface {
	Call(ctx context.Context, opts ...grpc.CallOption) (Data, error)
}

// ExecuteContractCaller is used to perform an execute contract call.
type ExecuteContractCaller interface {
	SendActionCaller

	SetGasPrice(*big.Int) ExecuteContractCaller
	SetGasLimit(uint64) ExecuteContractCaller
	SetAmount(*big.Int) ExecuteContractCaller
}

// Contract allows to read or execute on this contract's methods.
type Contract interface {
	ReadOnlyContract

	Execute(method string, args ...interface{}) ExecuteContractCaller
}

// ReadOnlyContract allows to read on this contract's methods.
type ReadOnlyContract interface {
	Read(method string, args ...interface{}) ReadContractCaller
}
