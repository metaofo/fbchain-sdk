package feesplit

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/client/context"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	"github.com/FiboChain/fbc/x/feesplit/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
)

const (
	// Amino names
	registerFeeSplitName = "fbexchain/MsgRegisterFeeSplit"
	updateFeeSplitName   = "fbexchain/MsgUpdateFeeSplit"
	cancelFeeSplitName   = "fbexchain/MsgCancelFeeSplit"
)

var (
	_ gosdktypes.Module = (*feesplitClient)(nil)
)

type feesplitClient struct {
	gosdktypes.BaseClient
	context.CLIContext
}

func (ibc feesplitClient) RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(types.MsgRegisterFeeSplit{}, registerFeeSplitName, nil)
	cdc.RegisterConcrete(types.MsgUpdateFeeSplit{}, updateFeeSplitName, nil)
	cdc.RegisterConcrete(types.MsgCancelFeeSplit{}, cancelFeeSplitName, nil)
}

// Name returns the module name
func (feesplitClient) Name() string {
	return types.ModuleName
}

// NewfeesplitClient creates a new instance of auth client as implement
func NewfeesplitClient(baseClient gosdktypes.BaseClient) feesplitClient {
	clientCtx := context.NewCLIContext().WithNodeURI(baseClient.GetConfig().NodeURI)
	return feesplitClient{baseClient, clientCtx}
}
