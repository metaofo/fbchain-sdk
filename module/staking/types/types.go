package types

import (
	stakingcli "github.com/FiboChain/fbc/x/staking/client/cli"
	stakingtypes "github.com/FiboChain/fbc/x/staking/types"
)

// const
const (
	ModuleName = stakingtypes.ModuleName
)

type (
	Validator         = stakingtypes.Validator
	DelegatorResponse = stakingcli.DelegatorResponse
)
