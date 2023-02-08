package types

import (
	stakingcli "github.com/zhengjianfeng1103/fbc/x/staking/client/cli"
	stakingtypes "github.com/zhengjianfeng1103/fbc/x/staking/types"
)

// const
const (
	ModuleName = stakingtypes.ModuleName
)

type (
	Validator         = stakingtypes.Validator
	DelegatorResponse = stakingcli.DelegatorResponse
)
