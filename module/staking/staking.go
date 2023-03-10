package staking

import (
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/staking/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	"github.com/zhengjianfeng1103/fbc/x/staking"
)

var _ gosdktypes.Module = (*stakingClient)(nil)

type stakingClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in staking module
func (sc stakingClient) RegisterCodec(cdc *codec.Codec) {
	staking.RegisterCodec(cdc)
}

// Name returns the module name
func (stakingClient) Name() string {
	return types.ModuleName
}

// NewStakingClient creates a new instance of staking client as implement
func NewStakingClient(baseClient gosdktypes.BaseClient) exposed.Staking {
	return stakingClient{baseClient}
}
