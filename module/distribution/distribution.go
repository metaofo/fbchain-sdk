package distribution

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	"github.com/FiboChain/fbc/x/distribution"
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/distribution/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
)

var _ gosdktypes.Module = (*distrClient)(nil)

type distrClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in distribution module
func (dc distrClient) RegisterCodec(cdc *codec.Codec) {
	distribution.RegisterCodec(cdc)
}

// Name returns the module name
func (distrClient) Name() string {
	return types.ModuleName
}

// NewDistrClient creates a new instance of distribution client as implement
func NewDistrClient(baseClient gosdktypes.BaseClient) exposed.Distribution {
	return distrClient{baseClient}
}
