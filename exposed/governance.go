package exposed

import (
	"github.com/metaofo/fbchain-sdk/module/governance/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

// Governance shows the expected behavior for inner governance client
type Governance interface {
	gosdktypes.Module
	GovTx
	GovQuery
}

// GovTx shows the expected tx behavior for inner governance client
type GovTx interface {
	SubmitTextProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	SubmitParamsChangeProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	SubmitDelistProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	SubmitCommunityPoolSpendProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	SubmitManageWhiteListProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Deposit(fromInfo keys.Info, passWd, depositCoinsStr, memo string, proposalID, accNum, seqNum uint64) (sdk.TxResponse, error)
	Vote(fromInfo keys.Info, passWd, voteOption, memo string, proposalID, accNum, seqNum uint64) (sdk.TxResponse, error)
}

// GovQuery shows the expected query behavior for inner governance client
type GovQuery interface {
	QueryProposals(depositorAddrStr, voterAddrStr, status string, numLimit uint64) ([]types.Proposal, error)
}
