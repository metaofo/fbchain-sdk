package evm

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	evm "github.com/FiboChain/fbc/x/evm/types"
	"github.com/FiboChain/fbchain-sdk/exposed"
	"github.com/FiboChain/fbchain-sdk/module/evm/types"
	gosdktypes "github.com/FiboChain/fbchain-sdk/types"
)

var _ gosdktypes.Module = (*evmClient)(nil)

type evmClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in evm module
func (ec evmClient) RegisterCodec(cdc *codec.Codec) {
	evm.RegisterCodec(cdc)
}

// Name returns the module name
func (evmClient) Name() string {
	return types.ModuleName
}

// NewEvmClient creates a new instance of evm client as implement
func NewEvmClient(baseClient gosdktypes.BaseClient) exposed.Evm {
	return evmClient{baseClient}
}