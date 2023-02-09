package main

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

func main() {
	var rpcClient *rpc.Client
	var err error

	rpcClient, err = rpc.Dial("http://8.210.109.185:8545")
	if err != nil {
		panic(err)
	}

	//txHash := "0x60b92276931041fd7d27ea52f9d169e1ff6658486e0471275a27f6118024b956"
	txHash := "0xe1ec7d2526435e0b7542d4ec024de7dd33c964bdadf7f01a3382960cacc2b8ae"

	//0xe1ec7d2526435e0b7542d4ec024de7dd33c964bdadf7f01a3382960cacc2b8ae

	var result interface{}
	err = rpcClient.Call(&result, "debug_traceTransaction", txHash, map[string]string{
		"tracer": "callTracer",
	})
	if err != nil {
		log.Println("debug_traceTransaction err:", err)
	}

	marshal, err := json.Marshal(result)
	if err != nil {
		log.Println("json.Marshal err:", err)
	}

	log.Println("result: ", string(marshal))

}
