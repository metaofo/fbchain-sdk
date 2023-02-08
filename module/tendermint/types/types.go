package types

import (
	ctypes "github.com/zhengjianfeng1103/fbc/libs/tendermint/rpc/core/types"
	tmtypes "github.com/zhengjianfeng1103/fbc/libs/tendermint/types"
)

// const
const (
	ModuleName  = "tendermint"
	EventFormat = "{eventType}.{eventAttribute}={value}"
)

type (
	Block              = tmtypes.Block
	ResultBlockResults = ctypes.ResultBlockResults
	ResultCommit       = ctypes.ResultCommit
	ResultValidators   = ctypes.ResultValidators
	ResultTx           = ctypes.ResultTx
	ResultTxSearch     = ctypes.ResultTxSearch
)
