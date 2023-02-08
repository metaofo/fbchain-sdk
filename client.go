package gosdk

import (
	"fmt"
	"github.com/metaofo/fbchain-sdk/module/feesplit"
	"github.com/metaofo/fbchain-sdk/module/ibc"
	ibcTypes "github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/apps/transfer/types"
	feesplitTypes "github.com/zhengjianfeng1103/fbc/x/feesplit/types"

	"github.com/metaofo/fbchain-sdk/exposed"
	"github.com/metaofo/fbchain-sdk/module"
	"github.com/metaofo/fbchain-sdk/module/ammswap"
	ammswaptypes "github.com/metaofo/fbchain-sdk/module/ammswap/types"
	"github.com/metaofo/fbchain-sdk/module/auth"
	authtypes "github.com/metaofo/fbchain-sdk/module/auth/types"
	"github.com/metaofo/fbchain-sdk/module/dex"
	dextypes "github.com/metaofo/fbchain-sdk/module/dex/types"
	"github.com/metaofo/fbchain-sdk/module/distribution"
	distrtypes "github.com/metaofo/fbchain-sdk/module/distribution/types"
	"github.com/metaofo/fbchain-sdk/module/evm"
	evmtypes "github.com/metaofo/fbchain-sdk/module/evm/types"
	"github.com/metaofo/fbchain-sdk/module/farm"
	"github.com/metaofo/fbchain-sdk/module/governance"
	govtypes "github.com/metaofo/fbchain-sdk/module/governance/types"
	"github.com/metaofo/fbchain-sdk/module/order"
	ordertypes "github.com/metaofo/fbchain-sdk/module/order/types"
	"github.com/metaofo/fbchain-sdk/module/slashing"
	slashingtypes "github.com/metaofo/fbchain-sdk/module/slashing/types"
	"github.com/metaofo/fbchain-sdk/module/staking"
	stakingtypes "github.com/metaofo/fbchain-sdk/module/staking/types"
	"github.com/metaofo/fbchain-sdk/module/tendermint"
	tmtypes "github.com/metaofo/fbchain-sdk/module/tendermint/types"
	"github.com/metaofo/fbchain-sdk/module/token"
	tokentypes "github.com/metaofo/fbchain-sdk/module/token/types"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	farmtypes "github.com/zhengjianfeng1103/fbc/x/farm/types"
)

// Client - structure of the main client of FbChain GoSDK
type Client struct {
	config  gosdktypes.ClientConfig
	cdc     *codec.Codec
	modules map[string]gosdktypes.Module
}

// NewClient creates a new instance of Client
func NewClient(config gosdktypes.ClientConfig) Client {
	cdc := gosdktypes.NewCodec()
	pClient := &Client{
		config:  config,
		cdc:     cdc,
		modules: make(map[string]gosdktypes.Module),
	}
	pBaseClient := module.NewBaseClient(cdc, &pClient.config)

	pClient.registerModule(
		ammswap.NewAmmSwapClient(pBaseClient),
		auth.NewAuthClient(pBaseClient),
		dex.NewDexClient(pBaseClient),
		distribution.NewDistrClient(pBaseClient),
		evm.NewEvmClient(pBaseClient),
		farm.NewFarmClient(pBaseClient),
		governance.NewGovClient(pBaseClient),
		order.NewOrderClient(pBaseClient),
		staking.NewStakingClient(pBaseClient),
		slashing.NewSlashingClient(pBaseClient),
		token.NewTokenClient(pBaseClient),
		tendermint.NewTendermintClient(pBaseClient),
		ibc.NewIbcClient(pBaseClient),
		feesplit.NewfeesplitClient(pBaseClient),
	)

	return *pClient
}

func (cli *Client) registerModule(mods ...gosdktypes.Module) {
	for _, mod := range mods {
		moduleName := mod.Name()
		if _, ok := cli.modules[moduleName]; ok {
			panic(fmt.Sprintf("duplicated module: %s", moduleName))
		}
		// register codec by each module
		mod.RegisterCodec(cli.cdc)
		cli.modules[moduleName] = mod
	}
	gosdktypes.RegisterBasicCodec(cli.cdc)
	cli.cdc.Seal()
}

// GetConfig returns the client config
func (cli *Client) GetConfig() gosdktypes.ClientConfig {
	return cli.config
}

// nolint
func (cli *Client) AmmSwap() exposed.AmmSwap {
	return cli.modules[ammswaptypes.ModuleName].(exposed.AmmSwap)
}
func (cli *Client) Auth() exposed.Auth {
	return cli.modules[authtypes.ModuleName].(exposed.Auth)
}
func (cli *Client) Dex() exposed.Dex {
	return cli.modules[dextypes.ModuleName].(exposed.Dex)
}
func (cli *Client) Distribution() exposed.Distribution {
	return cli.modules[distrtypes.ModuleName].(exposed.Distribution)
}
func (cli *Client) Evm() exposed.Evm {
	return cli.modules[evmtypes.ModuleName].(exposed.Evm)
}
func (cli *Client) Farm() exposed.Farm {
	return cli.modules[farmtypes.ModuleName].(exposed.Farm)
}
func (cli *Client) Governance() exposed.Governance {
	return cli.modules[govtypes.ModuleName].(exposed.Governance)
}
func (cli *Client) Order() exposed.Order {
	return cli.modules[ordertypes.ModuleName].(exposed.Order)
}
func (cli *Client) Slashing() exposed.Slashing {
	return cli.modules[slashingtypes.ModuleName].(exposed.Slashing)
}
func (cli *Client) Staking() exposed.Staking {
	return cli.modules[stakingtypes.ModuleName].(exposed.Staking)
}
func (cli *Client) Tendermint() exposed.Tendermint {
	return cli.modules[tmtypes.ModuleName].(exposed.Tendermint)
}
func (cli *Client) Token() exposed.Token {
	return cli.modules[tokentypes.ModuleName].(exposed.Token)
}

func (cli *Client) Ibc() exposed.Ibc {
	return cli.modules[ibcTypes.ModuleName].(exposed.Ibc)
}

func (cli *Client) Feesplit() exposed.Feesplit {
	return cli.modules[feesplitTypes.ModuleName].(exposed.Feesplit)
}
