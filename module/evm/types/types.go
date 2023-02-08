package types

import (
	"math/big"

	apptypes "github.com/zhengjianfeng1103/fbc/app/types"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	evmtypes "github.com/zhengjianfeng1103/fbc/x/evm/types"
)

// const
const (
	ModuleName      = evmtypes.ModuleName
	defaultGasPrice = "0.000000001"
)

type (
	QueryResCode    = evmtypes.QueryResCode
	QueryResStorage = evmtypes.QueryResStorage
)

var (
	DefaultGasPrice    = sdk.MustNewDecFromStr(defaultGasPrice).BigInt()
	DefaultRPCGasLimit = big.NewInt(apptypes.DefaultRPCGasLimit)
)
