package auth

import (
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/auth/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth"
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
