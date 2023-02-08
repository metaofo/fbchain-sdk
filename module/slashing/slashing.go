package slashing

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	"github.com/FiboChain/fbc/x/slashing"
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/slashing/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
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
