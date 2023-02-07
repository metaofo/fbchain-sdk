package types

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth/exported"
)

// const
const (
	ModuleName = auth.ModuleName
)

type (
	Account = exported.Account
)
