package evm

import (
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/evm/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	evm "github.com/zhengjianfeng1103/fbc/x/evm/types"
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
