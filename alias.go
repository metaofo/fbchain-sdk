package gosdk

import (
	sdk "github.com/FiboChain/fbc/libs/cosmos-sdk/types"
	farm "github.com/FiboChain/fbc/x/farm/types"
	ammswap "github.com/metaofo/fbchain-sdk/module/ammswap/types"
	auth "github.com/metaofo/fbchain-sdk/module/auth/types"
	dex "github.com/metaofo/fbchain-sdk/module/dex/types"
	evm "github.com/metaofo/fbchain-sdk/module/evm/types"
	governance "github.com/metaofo/fbchain-sdk/module/governance/types"
	order "github.com/metaofo/fbchain-sdk/module/order/types"
	staking "github.com/metaofo/fbchain-sdk/module/staking/types"
	tendermint "github.com/metaofo/fbchain-sdk/module/tendermint/types"
	token "github.com/metaofo/fbchain-sdk/module/token/types"
	"github.com/metaofo/fbchain-sdk/types"
)

// const
const (
	BroadcastSync  = types.BroadcastSync
	BroadcastAsync = types.BroadcastAsync
	BroadcastBlock = types.BroadcastBlock

	// vote for the proposal
	VoteYes        = "yes"
	VoteAbstain    = "abstain"
	VoteNo         = "no"
	VoteNoWithVeto = "no_with_veto"
)

var (
	// NewClientConfig gives an easy way for the callers to set client config
	NewClientConfig = types.NewClientConfig
)

// nolint
type (
	TxResponse = sdk.TxResponse
	// ammswap
	SwapTokenPair = ammswap.SwapTokenPair
	// auth
	Account = auth.Account
	// staking
	Validator         = staking.Validator
	DelegatorResponse = staking.DelegatorResponse
	// token
	TokenResp = token.TokenResp
	// dex
	TokenPair = dex.TokenPair
	// order
	BookRes     = order.BookRes
	OrderDetail = order.OrderDetail
	// tendermint
	Block            = tendermint.Block
	BlockResults     = tendermint.ResultBlockResults
	ResultCommit     = tendermint.ResultCommit
	ResultValidators = tendermint.ResultValidators
	ResultTx         = tendermint.ResultTx
	ResultTxSearch   = tendermint.ResultTxSearch
	// governance
	Proposal = governance.Proposal
	// farm
	FarmPool = farm.FarmPool
	LockInfo = farm.LockInfo
	// evm
	QueryResCode    = evm.QueryResCode
	QueryResStorage = evm.QueryResStorage
)
