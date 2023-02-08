package farm

import (
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/farm/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	farm "github.com/zhengjianfeng1103/fbc/x/farm/types"
)

type farmClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in farm module
func (fc farmClient) RegisterCodec(cdc *codec.Codec) {
	farm.RegisterCodec(cdc)
}

// Name returns the module name
func (farmClient) Name() string {
	return types.ModuleName
}

// NewFarmClient creates a new instance of farm client as implement
func NewFarmClient(baseClient gosdktypes.BaseClient) exposed.Farm {
	return farmClient{baseClient}
}
