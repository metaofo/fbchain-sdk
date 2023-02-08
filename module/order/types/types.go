package types

import (
	orderkeeper "github.com/zhengjianfeng1103/fbc/x/order/keeper"
	ordertypes "github.com/zhengjianfeng1103/fbc/x/order/types"
)

// const
const (
	ModuleName = ordertypes.ModuleName
)

type (
	BookRes     = orderkeeper.BookRes
	OrderDetail = ordertypes.Order
)
