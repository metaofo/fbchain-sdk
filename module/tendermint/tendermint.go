package tendermint

import (
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/tendermint/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
)

var _ gosdktypes.Module = (*tendermintClient)(nil)

type tendermintClient struct {
	gosdktypes.BaseClient
}

// nolint
func (tendermintClient) RegisterCodec(_ *codec.Codec) {}

// Name returns the module name
func (tendermintClient) Name() string {
	return types.ModuleName
}

// NewTendermintClient creates a new instance of tendermint client as implement
func NewTendermintClient(baseClient gosdktypes.BaseClient) exposed.Tendermint {
	return tendermintClient{baseClient}
}
