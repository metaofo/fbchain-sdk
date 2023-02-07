package auth

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth"
	"github.com/FiboChain/fbchain-sdk/exposed"
	"github.com/FiboChain/fbchain-sdk/module/auth/types"
	gosdktypes "github.com/FiboChain/fbchain-sdk/types"
)

var _ gosdktypes.Module = (*authClient)(nil)

type authClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the struct type for auth module
func (authClient) RegisterCodec(cdc *codec.Codec) {
	auth.RegisterCodec(cdc)
}

// Name returns the module name
func (authClient) Name() string {
	return types.ModuleName
}

// NewAuthClient creates a new instance of auth client as implement
func NewAuthClient(baseClient gosdktypes.BaseClient) exposed.Auth {
	return authClient{baseClient}
}
