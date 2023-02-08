package types

import (
	cryptocodec "github.com/zhengjianfeng1103/fbc/app/crypto/ethsecp256k1"
	evmtypes "github.com/zhengjianfeng1103/fbc/app/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
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
