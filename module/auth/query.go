package auth

import (
	"errors"
	"fmt"

	sdk "github.com/FiboChain/fbc/libs/cosmos-sdk/types"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth"
	authtypes "github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth/types"
	"github.com/FiboChain/fbchain-sdk/module/auth/types"
	"github.com/FiboChain/fbchain-sdk/utils"
)

// QueryAccount gets the account info
func (ac authClient) QueryAccount(accAddrStr string) (account types.Account, err error) {
	accAddr, err := sdk.AccAddressFromBech32(accAddrStr)
	if err != nil {
		return account, errors.New("failed. accAddress converted from Bech32 error")
	}

	path := fmt.Sprintf("custom/%s/%s", auth.QuerierRoute, auth.QueryAccount)
	bytes, err := ac.GetCodec().MarshalJSON(authtypes.NewQueryAccountParams(accAddr))
	if err != nil {
		return account, utils.ErrClientQuery(err.Error())
	}

	res, _, err := ac.Query(path, bytes)
	if res == nil {
		return account, fmt.Errorf("failed. your account has no record on the chain. error: %s", err)
	}

	if err = ac.GetCodec().UnmarshalJSON(res, &account); err != nil {
		return account, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
