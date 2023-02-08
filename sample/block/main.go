package main

import (
	gosdk "github.com/metaofo/fbchain-sdk"
	"github.com/zhengjianfeng1103/fbc/app"
	"github.com/zhengjianfeng1103/fbc/app/codec"
	"github.com/zhengjianfeng1103/fbc/app/rpc/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/context"
	sdkcodec "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	client_types "github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/02-client/types"
	"github.com/zhengjianfeng1103/fbc/x/evm/watcher"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	tmtypes "github.com/zhengjianfeng1103/fbc/libs/tendermint/types"
	"log"
	"math/big"
)

const rpcURL = "tcp://8.210.109.185:26657"
const rpcURL2 = "http//8.210.109.185:8545"

// fbchaincli tx send kk 0xFd2955B33Fa3bE18b6ef3a90097F8a25F5E5FF85 0.01fibo --gas-prices 0.000001fibo
// fbchaincli query tx 267FA72CDC8321253CAA21EC0E6B1E63BE45147D28F6392AEB42D1EE1BB48176 --trust-node

// curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getAllTransactionResultsByBlock","params":["0x9970B8", "0x0", "0xA"],"id":1}' -H "Content-Type: application/json" http://8.210.109.185:8545
// "type":"0x1",
//cosmos-sdk/StdTx

// curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getAllTransactionResultsByBlock","params":["0x905270", "0x0", "0xA"],"id":1}' -H "Content-Type: application/json" http://8.210.109.185:8545
// "type":"0x0",

func main() {
	tmtypes.UnittestOnlySetMilestoneVenusHeight(1)

	config, err := gosdk.NewClientConfig(rpcURL, "fbc-1230", gosdk.BroadcastBlock, "0.01fibo", 200000,
		0, "")
	if err != nil {
		log.Fatal(err)
	}

	rpcClient := gosdk.NewClient(config)

	//10055864
	block, err := rpcClient.Tendermint().QueryBlock(10055864)
	if err != nil {
		log.Fatal(err)
	}

	interfaceReg := codec.MakeIBC(app.ModuleBasics)
	makeCodec := codec.MakeCodec(app.ModuleBasics)
	protoCdc := sdkcodec.NewProtoCodec(interfaceReg)
	proxy := sdkcodec.NewCodecProxy(protoCdc, makeCodec)
	cliContext := context.NewCLIContext().
		WithNodeURI(rpcURL).
		WithProxy(proxy).
		WithTrustNode(true)

	txs := block.Txs
	chainIdStr := block.ChainID
	chainId := client_types.ParseChainID(chainIdStr)
	log.Println("chainId: ", chainId)
	for idx, txRes := range txs {

		height := tmtypes.GetVenusHeight()
		log.Println("venus height: ", height)

		venus := tmtypes.HigherThanVenus(block.Height)
		log.Println("venus: ", venus)

		txHash := txRes.Hash(block.Height)

		realTx, err := types.RawTxToRealTx(cliContext, txRes,
			common.BytesToHash(block.Hash()), uint64(block.Height), uint64(idx+1))
		if err != nil {
			log.Fatal("RawTxToRealTx: ", err)
		}

		log.Println("type: ", realTx.GetType())
		log.Println("from: ", realTx.GetFrom())

		queryTx, err := cliContext.Client.Tx(txHash, false)
		if err != nil {
			log.Fatal("queryTx: ", err)
		}

		switch realTx.GetType() {
		case sdk.EvmTxType:
			receipt, err := types.RawTxResultToEthReceipt(big.NewInt(int64(chainId)), queryTx, realTx, common.BytesToHash(block.Hash()))
			if err != nil {
				log.Fatal("RawTxResultToEthReceipt: ", err)
			}

			log.Println("RawTxResultToEthReceipt: ", receipt)

		case sdk.StdTxType:
			res, err := watcher.RawTxResultToStdResponse(cliContext, queryTx, realTx, block.Time)
			if err != nil {
				log.Fatal("RawTxResultToStdResponse: ", err)
			}

			log.Printf("RawTxResultToStdResponse: %+v", res.Response.Response)
		default:
			log.Fatal("unknown tx type")
		}

	}

}

var web3Client *ethclient.Client

// rpc.enable-multi-call
func init() {
	web3Client, _ = ethclient.Dial(rpcURL2)

}
