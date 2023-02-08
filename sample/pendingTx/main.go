package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zhengjianfeng1103/fbc/app"
	"github.com/zhengjianfeng1103/fbc/app/codec"
	sdkContext "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/context"
	sdkcodec "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	"github.com/zhengjianfeng1103/fbc/x/evm/watcher"

	"log"
)

var wssClient *rpc.Client
var web3Client *ethclient.Client
var cliContext sdkContext.CLIContext

func init() {
	var err error
	wssClient, err = rpc.Dial("ws://8.210.109.185:8546")
	if err != nil {
		panic(err)
	}
	//0xef704c9f40465129aed7ed4b579471881c641169c22de8fade51d3651e9c0752

	//web3Client, err = ethclient.Dial("https://node.yiwos.com")
	web3Client, err = ethclient.Dial("http://8.210.109.185:8545")
	if err != nil {
		panic(err)
	}

	interfaceReg := codec.MakeIBC(app.ModuleBasics)
	makeCodec := codec.MakeCodec(app.ModuleBasics)
	protoCdc := sdkcodec.NewProtoCodec(interfaceReg)
	proxy := sdkcodec.NewCodecProxy(protoCdc, makeCodec)
	cliContext = sdkContext.NewCLIContext().
		WithNodeURI("tcp://8.210.109.185:26657").
		WithProxy(proxy).
		WithTrustNode(true)
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic err:", err)
		}
	}()

	ReconnectTimes := 0
Reconnect:

	ReconnectTimes++
	if ReconnectTimes > 1000 {
		log.Println("ReconnectTimes > 1000")
		return
	}
	pendingTxRequest := "newPendingTransactions"
	dataChannel := make(chan string)
	subscribe, err := wssClient.EthSubscribe(context.Background(), dataChannel, pendingTxRequest)
	if err != nil {
		log.Fatal("subscribe eth: " + err.Error())
	}

	log.Println("subscribe success, begin to listen data.....")

	for {
		select {
		case err = <-subscribe.Err():
			log.Println("subscribe err: " + err.Error())
			goto Reconnect
		case txHash := <-dataChannel:
			// do something
			log.Println("txHash:", txHash)

			tx, err := cliContext.Client.GetUnconfirmedTxByHash(common.HexToHash(txHash))
			if err != nil {
				log.Println("GetUnconfirmedTxByHash err:", err)
			}

			wTx, pending, err := web3Client.TransactionByHash(context.Background(), common.HexToHash(txHash))
			if err != nil {
				log.Println("TransactionByHash err:", err)
			}

			log.Println("wTx", wTx, "pending:", pending)

			var watchTxHash watcher.Transaction
			err = wssClient.Call(&watchTxHash, "eth_getTransactionByHash", txHash)
			if err != nil {
				log.Println("eth_getTransactionByHash err:", err)
			}

			log.Println("watchTxHash:", watchTxHash, "tx", tx)

			if tx == nil || wTx == nil || watchTxHash.Hash.String() == "" {
				log.Println("tx is nil")
				continue
			}

			//log.Println("pending:", tx.Height, tx.TxResult.Code, tx.TxResult.Log)
		}
	}

}
