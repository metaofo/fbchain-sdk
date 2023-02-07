package exposed

import (
	"github.com/FiboChain/fbchain-sdk/module/auth/types"
	gosdktypes "github.com/FiboChain/fbchain-sdk/types"
)

// Auth shows the expected behavior for inner auth client
type Auth interface {
	gosdktypes.Module
	AuthQuery
}

// AuthQuery shows the expected query behavior for inner auth client
type AuthQuery interface {
	QueryAccount(accAddrStr string) (types.Account, error)
}
