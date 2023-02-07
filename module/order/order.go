package order

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	"github.com/FiboChain/fbc/x/order"
	"github.com/FiboChain/fbchain-sdk/exposed"
	"github.com/FiboChain/fbchain-sdk/module/order/types"
	gosdktypes "github.com/FiboChain/fbchain-sdk/types"
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
