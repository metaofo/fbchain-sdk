package exposed

import (
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

// Slashing shows the expected behavior for inner slashing client
type Slashing interface {
	gosdktypes.Module
	SlashingTx
}

// SlashingTx shows the expected tx behavior for inner slashing client
type SlashingTx interface {
	Unjail(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}
