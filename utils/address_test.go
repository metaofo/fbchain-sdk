package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

const (
	defaultValAddr            = "fbvaloper1qj5c07sm6jetjz8f509qtrxgh4psxkv3m2wy6x"
	accAddrWithFbcChainPrefix = "fbexchain1qj5c07sm6jetjz8f509qtrxgh4psxkv32x0qas"
	valAddrWithFbcChainPrefix = "fbexchainvaloper1qj5c07sm6jetjz8f509qtrxgh4psxkv3tzllpj"
)

func TestAccAddrPrefixConvert(t *testing.T) {
	dstAccAddrStr, err := AccAddrPrefixConvert("ex", defaultAddr, "fbexchain")
	require.NoError(t, err)
	require.Equal(t, accAddrWithFbcChainPrefix, dstAccAddrStr)

	// error check
	// wrong source prefix
	_, err = AccAddrPrefixConvert("exx", defaultAddr, "fbexchain")
	require.Error(t, err)

	// wrong source account address
	_, err = AccAddrPrefixConvert("ex", defaultAddr+"a", "fbexchain")
	require.Error(t, err)

	// wrong destination account address
	dstAccAddrStr, err = AccAddrPrefixConvert("ex", defaultAddr, "fbexchainx")
	require.NoError(t, err)
	require.NotEqual(t, accAddrWithFbcChainPrefix, dstAccAddrStr)

	// recover the account prefix
	sdk.GetConfig().SetBech32PrefixForAccount("ex", "expub")
}

func TestValAddrPrefixConvert(t *testing.T) {
	dstValAddrStr, err := ValAddrPrefixConvert("fbvaloper", defaultValAddr, "fbexchainvaloper")
	require.NoError(t, err)
	require.Equal(t, valAddrWithFbcChainPrefix, dstValAddrStr)

	// error check
	// wrong source prefix
	_, err = ValAddrPrefixConvert("exxvaloper", defaultValAddr, "fbexchainvaloper")
	require.Error(t, err)

	// wrong source validator address
	_, err = ValAddrPrefixConvert("fbvaloper", defaultValAddr+"a", "fbexchainvaloper")
	require.Error(t, err)

	// wrong destination validator address
	dstValAddrStr, err = ValAddrPrefixConvert("fbvaloper", defaultValAddr, "fbexchainxvaloper")
	require.NoError(t, err)
	require.NotEqual(t, valAddrWithFbcChainPrefix, dstValAddrStr)

	// recover the validator prefix
	sdk.GetConfig().SetBech32PrefixForValidator("fbvaloper", "fbvaloperpub")
}
