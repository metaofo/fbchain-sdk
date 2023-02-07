package farm

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	farm "github.com/FiboChain/fbc/x/farm/types"
	"github.com/FiboChain/fbchain-sdk/exposed"
	"github.com/FiboChain/fbchain-sdk/module/farm/types"
	gosdktypes "github.com/FiboChain/fbchain-sdk/types"
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
