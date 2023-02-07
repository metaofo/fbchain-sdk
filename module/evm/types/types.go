package types

import (
	"math/big"

	apptypes "github.com/FiboChain/fbc/app/types"
	sdk "github.com/FiboChain/fbc/libs/cosmos-sdk/types"
	evmtypes "github.com/FiboChain/fbc/x/evm/types"
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
