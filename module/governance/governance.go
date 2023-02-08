package governance

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	"github.com/FiboChain/fbc/x/gov"
	paramstypes "github.com/FiboChain/fbc/x/params/types"
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/governance/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
)

var _ gosdktypes.Module = (*govClient)(nil)

type govClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in governance module
func (gc govClient) RegisterCodec(cdc *codec.Codec) {
	gov.RegisterCodec(cdc)
	cdc.RegisterConcrete(paramstypes.ParameterChangeProposal{}, "fbexchain/params/ParameterChangeProposal", nil)
}

// Name returns the module name
func (gc govClient) Name() string {
	return types.ModuleName
}

// NewGovClient creates a new instance of governance client as implement
func NewGovClient(baseClient gosdktypes.BaseClient) exposed.Governance {
	return govClient{baseClient}
}
