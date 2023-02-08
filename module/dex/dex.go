package dex

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	"github.com/FiboChain/fbc/x/dex"
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/dex/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
)

var _ gosdktypes.Module = (*dexClient)(nil)

type dexClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in dex module
func (dexClient) RegisterCodec(cdc *codec.Codec) {
	dex.RegisterCodec(cdc)
}

// Name returns the module name
func (dexClient) Name() string {
	return types.ModuleName
}

// NewDexClient creates a new instance of dex client as implement
func NewDexClient(baseClient gosdktypes.BaseClient) exposed.Dex {
	return dexClient{baseClient}
}
