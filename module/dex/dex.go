package dex

import (
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/dex/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	"github.com/zhengjianfeng1103/fbc/x/dex"
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
