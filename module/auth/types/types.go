package types

import (
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth/exported"
)

// const
const (
	ModuleName = auth.ModuleName
)

type (
	Account = exported.Account
)
