package order

import (
	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module/order/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	"github.com/zhengjianfeng1103/fbc/x/order"
)

var _ gosdktypes.Module = (*orderClient)(nil)

type orderClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in order module
func (orderClient) RegisterCodec(cdc *codec.Codec) {
	order.RegisterCodec(cdc)
}

// Name returns the module name
func (orderClient) Name() string {
	return types.ModuleName
}

// NewOrderClient creates a new instance of order client as implement
func NewOrderClient(baseClient gosdktypes.BaseClient) exposed.Order {
	return orderClient{baseClient}
}
