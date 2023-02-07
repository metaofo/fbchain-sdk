package slashing

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/FiboChain/fbc/libs/cosmos-sdk/types"
	"github.com/FiboChain/fbc/x/slashing"
	"github.com/FiboChain/fbchain-sdk/types/params"
)

// Unjail unjails the own validator which was jailed by slashing module
func (sc slashingClient) Unjail(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (resp sdk.TxResponse,
	err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := slashing.NewMsgUnjail(sdk.ValAddress(fromInfo.GetAddress()))
	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}