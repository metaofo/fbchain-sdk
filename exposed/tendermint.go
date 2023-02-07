package exposed

import (
	ctypes "github.com/FiboChain/fbc/libs/tendermint/rpc/core/types"
	"github.com/FiboChain/fbchain-sdk/module/tendermint/types"
	gosdktypes "github.com/FiboChain/fbchain-sdk/types"
)

// Tendermint shows the expected behavior for inner tendermint client
type Tendermint interface {
	gosdktypes.Module
	TendermintQuery
}

// TendermintQuery shows the expected query behavior for inner tendermint client
type TendermintQuery interface {
	QueryBlock(height int64) (*types.Block, error)
	QueryBlockResults(height int64) (*types.ResultBlockResults, error)
	QueryCommitResult(height int64) (*types.ResultCommit, error)
	QueryValidatorsResult(height int64) (*types.ResultValidators, error)
	QueryTxResult(hashHexStr string, prove bool) (*types.ResultTx, error)
	// QueryTxsByEvents assumes the node to query a truth teller
	QueryTxsByEvents(eventsStr string, page, limit int) (*ctypes.ResultTxSearch, error)
	QueryStatus() (*ctypes.ResultStatus, error)
}
