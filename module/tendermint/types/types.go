package types

import (
	ctypes "github.com/FiboChain/fbc/libs/tendermint/rpc/core/types"
	tmtypes "github.com/FiboChain/fbc/libs/tendermint/types"
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
