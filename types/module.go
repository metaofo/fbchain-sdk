package types

import "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"

// Module shows the expected behaviour of each module in FbChain GoSDK
type Module interface {
	RegisterCodec(cdc *codec.Codec)
	Name() string
}
