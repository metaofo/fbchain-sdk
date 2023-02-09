package main

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/rpc"
	gosdk "github.com/metaofo/fbchain-sdk"
	"github.com/zhengjianfeng1103/fbc/app"
	"github.com/zhengjianfeng1103/fbc/app/codec"
	"github.com/zhengjianfeng1103/fbc/app/rpc/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/context"
	sdkcodec "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	client_types "github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/02-client/types"
	"github.com/zhengjianfeng1103/fbc/x/evm/watcher"
	"time"

	"github.com/ethereum/go-ethereum/common"
	tmtypes "github.com/zhengjianfeng1103/fbc/libs/tendermint/types"
	"log"
	"math/big"
)

const rpcURL = "tcp://8.210.109.185:26657"
const rpcURL2 = "http://8.210.109.185:8545"

// fbchaincli tx send kk 0xFd2955B33Fa3bE18b6ef3a90097F8a25F5E5FF85 0.01fibo --gas-prices 0.000001fibo
// fbchaincli query tx 267FA72CDC8321253CAA21EC0E6B1E63BE45147D28F6392AEB42D1EE1BB48176 --trust-node

// curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getAllTransactionResultsByBlock","params":["0x9970B8", "0x0", "0xA"],"id":1}' -H "Content-Type: application/json" http://8.210.109.185:8545
// "type":"0x1",
//cosmos-sdk/StdTx

// curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getAllTransactionResultsByBlock","params":["0x905270", "0x0", "0xA"],"id":1}' -H "Content-Type: application/json" http://8.210.109.185:8545
// "type":"0x0",

type CallTracerResult struct {
	JsonRpc    string `json:"jsonrpc"`
	Id         int    `json:"id"`
	CallTracer struct {
		Calls []struct {
			From   string `json:"from"`
			Input  string `json:"input"`
			Output string `json:"output"`
			To     string `json:"to"`
			Type   string `json:"type"`
			Value  string `json:"value"`
		} `json:"calls"`
		From    string `json:"from"`
		Gas     string `json:"gas"`
		GasUsed string `json:"gasUsed"`
		Input   string `json:"input"`
		Output  string `json:"output"`
		Time    string `json:"time"`
		To      string `json:"to"`
		Type    string `json:"type"`
		Value   string `json:"value"`
	} `json:"result"`
}

type CallTracer struct {
	From    string `json:"from"`
	Gas     string `json:"gas"`
	GasUsed string `json:"gasUsed"`
	Input   string `json:"input"`
	Output  string `json:"output"`
	Time    string `json:"time"`
	To      string `json:"to"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

func main() {
	tmtypes.UnittestOnlySetMilestoneVenusHeight(1)

	config, err := gosdk.NewClientConfig(rpcURL, "fbc-1230", gosdk.BroadcastBlock, "0.01fibo", 200000,
		0, "")
	if err != nil {
		log.Fatal(err)
	}

	rpcClient := gosdk.NewClient(config)

	ethRpcClient, err := rpc.Dial(rpcURL2)
	if err != nil {
		panic(err)
	}

	//10055864
	interfaceReg := codec.MakeIBC(app.ModuleBasics)
	makeCodec := codec.MakeCodec(app.ModuleBasics)
	protoCdc := sdkcodec.NewProtoCodec(interfaceReg)
	proxy := sdkcodec.NewCodecProxy(protoCdc, makeCodec)
	cliContext := context.NewCLIContext().
		WithNodeURI(rpcURL).
		WithProxy(proxy).
		WithTrustNode(true)

	for startBlock := 10055864; startBlock < 10055864+10000; startBlock++ {
		time.Sleep(100 * time.Millisecond)
		block, err := rpcClient.Tendermint().QueryBlock(int64(startBlock))
		if err != nil {
			log.Fatal(err)
		}

		txs := block.Txs
		chainIdStr := block.ChainID
		chainId := client_types.ParseChainID(chainIdStr)
		for idx, txRes := range txs {
			txHash := txRes.Hash(block.Height)

			realTx, err := types.RawTxToRealTx(cliContext, txRes,
				common.BytesToHash(block.Hash()), uint64(block.Height), uint64(idx+1))
			if err != nil {
				log.Fatal("RawTxToRealTx: ", err)
			}
			log.Println("=====================tx START=======================")
			log.Println("type: ", realTx.GetType())
			log.Println("from: ", realTx.GetFrom())
			log.Println("hash: ", common.BytesToHash(realTx.TxHash()).String())

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

				var callTracerResp map[string]interface{}
				err = ethRpcClient.Call(&callTracerResp, "debug_traceTransaction", receipt.EthTx.Hash, map[string]string{
					"tracer": "callTracer",
				})
				if err != nil {
					log.Fatal("debug_traceTransaction: ", err)
				}

				if receipt.EthTx.To == nil {
					log.Fatal("receipt.EthTx.To is nil")
				}

				var codeAt interface{}
				err = ethRpcClient.Call(&codeAt, "eth_getCode", receipt.EthTx.To.String(), "latest")
				if err != nil {
					log.Fatal("eth_getCode: ", err)
				}

				marshal, err := json.Marshal(callTracerResp)
				if err != nil {
					log.Fatal("json.Marshal: ", err)
				}

				if codeAt.(string) == "0x" {
					var callTracer CallTracer
					err = json.Unmarshal(marshal, &callTracer)
					if err != nil {
						log.Fatal("json.Unmarshal: ", err)
					}

					log.Printf("callTracer: %+v \n", callTracer)
				} else {
					var callTracerResult CallTracerResult
					err = json.Unmarshal(marshal, &callTracerResult)
					if err != nil {
						log.Fatal("json.Unmarshal: ", err)
					}

					log.Printf("callTracerResult: %+v \n", callTracerResult)
				}

			case sdk.StdTxType:
				res, err := watcher.RawTxResultToStdResponse(cliContext, queryTx, realTx, block.Time)
				if err != nil {
					log.Fatal("RawTxResultToStdResponse: ", err)
				}

				log.Printf("RawTxResultToStdResponse: %+v", res.Response.Response)
			default:
				log.Fatal("unknown tx type")
			}

			log.Println("=====================tx END=======================")

		}

	}

}
