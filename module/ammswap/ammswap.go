package ammswap

import (
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/ammswap/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	"github.com/zhengjianfeng1103/fbc/x/ammswap"
)

var _ gosdktypes.Module = (*ammswapClient)(nil)

type ammswapClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in ammswap module
func (ac ammswapClient) RegisterCodec(cdc *codec.Codec) {
	ammswap.RegisterCodec(cdc)
}

// Name returns the module name
func (ammswapClient) Name() string {
	return types.ModuleName
}

// NewAmmSwapClient creates a new instance of ammswap client as implement
func NewAmmSwapClient(baseClient gosdktypes.BaseClient) exposed.AmmSwap {
	return ammswapClient{baseClient}
}
