package slashing

import (
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/slashing/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	"github.com/zhengjianfeng1103/fbc/x/slashing"
)

var _ gosdktypes.Module = (*slashingClient)(nil)

type slashingClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in slashing module
func (sc slashingClient) RegisterCodec(cdc *codec.Codec) {
	slashing.RegisterCodec(cdc)
}

// Name returns the module name
func (slashingClient) Name() string {
	return types.ModuleName
}

// NewSlashingClient creates a new instance of slashing client as implement
func NewSlashingClient(baseClient gosdktypes.BaseClient) exposed.Slashing {
	return slashingClient{baseClient}
}
