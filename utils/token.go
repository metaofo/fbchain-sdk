package utils

import (
	"errors"
	"strings"

	sdk "github.com/FiboChain/fbc/libs/cosmos-sdk/types"
	tokentypes "github.com/FiboChain/fbc/x/token/types"
	"github.com/FiboChain/fbchain-sdk/module/token/types"
)

// ParseTransfersStr parses the whole multi-send info string into TransferUnit
// Example:
// `addr1 1fibo
// 	addr2 2fibo`
func ParseTransfersStr(str string) ([]types.TransferUnit, error) {
	strs := strings.Split(strings.TrimSpace(str), "\n")
	transLen := len(strs)
	transfers := make([]types.TransferUnit, transLen)

	for i := 0; i < transLen; i++ {
		s := strings.Split(strs[i], " ")
		if len(s) != 2 {
			return nil, errors.New("invalid text to parse")
		}
		addrStr, coinStr := s[0], s[1]

		to, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return nil, err
		}

		coins, err := sdk.ParseDecCoins(coinStr)
		if err != nil {
			return nil, err
		}

		transfers[i] = tokentypes.TransferUnit{
			To:    to,
			Coins: coins,
		}
	}

	return transfers, nil
}