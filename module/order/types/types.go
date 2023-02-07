package types

import (
	orderkeeper "github.com/FiboChain/fbc/x/order/keeper"
	ordertypes "github.com/FiboChain/fbc/x/order/types"
)

// const
const (
	ModuleName = ordertypes.ModuleName
)

type (
	BookRes     = orderkeeper.BookRes
	OrderDetail = ordertypes.Order
)
