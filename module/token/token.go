package token

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	"github.com/FiboChain/fbc/x/token"
	"github.com/FiboChain/fbchain-sdk/exposed"
	"github.com/FiboChain/fbchain-sdk/module/token/types"
	gosdktypes "github.com/FiboChain/fbchain-sdk/types"
)

var _ gosdktypes.Module = (*tokenClient)(nil)

type tokenClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in token module
func (tokenClient) RegisterCodec(cdc *codec.Codec) {
	token.RegisterCodec(cdc)
}

// Name returns the module name
func (tokenClient) Name() string {
	return types.ModuleName
}

// NewTokenClient creates a new instance of token client as implement
func NewTokenClient(baseClient gosdktypes.BaseClient) exposed.Token {
	return tokenClient{baseClient}
}
