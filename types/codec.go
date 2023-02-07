package types

import (
	cryptocodec "github.com/FiboChain/fbc/app/crypto/ethsecp256k1"
	evmtypes "github.com/FiboChain/fbc/app/types"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	sdk "github.com/FiboChain/fbc/libs/cosmos-sdk/types"
)

// NewCodec creates a new instance of codec only for gosdk
func NewCodec() *codec.Codec {
	return codec.New()
}

// RegisterBasicCodec registers the basic data types for gosdk codec
func RegisterBasicCodec(cdc *codec.Codec) {
	sdk.RegisterCodec(cdc)
	cryptocodec.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	evmtypes.RegisterCodec(cdc)
}
