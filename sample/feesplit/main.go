package main

import (
	"fmt"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/types/query"
	"github.com/FiboChain/fbchain-sdk/utils"
	"log"

	gosdk "github.com/FiboChain/fbchain-sdk"
)

const (
	// TODO: link to mainnet of FbChain later
	rpcURL = "tcp://52.199.88.250:26657"
	// user's name
	name = "alice"
	// user's mnemonic
	mnemonic = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	// user's password
	passWd = "12345678"
	// target address
	addr     = "fb1qj5c07sm6jetjz8f509qtrxgh4psxkv3htfcuj"
	baseCoin = "fibo"

	privateKey = ""
)

func main() {
	//-------------------- 1. preparation --------------------//
	// NOTE: either of the both ways below to pay fees is available

	// WAY 1: create a client config with fixed fees
	//config, err := gosdk.NewClientConfig(rpcURL, "fbc-64", gosdk.BroadcastBlock, "0.01fibo", 200000,
	//	0, "")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// WAY 2: alternative client config with the fees by auto gas calculation
	config, err := gosdk.NewClientConfig(rpcURL, "fbc-64", gosdk.BroadcastBlock, "", 200000,
		1.1, "0.000000001fibo")
	if err != nil {
		log.Fatal(err)
	}

	cli := gosdk.NewClient(config)

	// create an account with your own privateKey name and password
	fromInfo, err := utils.CreateAccountWithPrivateKey(privateKey, "admin", passWd)
	if err != nil {
		log.Fatal(err)
	}

	accInfo, err := cli.Auth().QueryAccount(fromInfo.GetAddress().String())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(accInfo)

	seq := accInfo.GetSequence()
	fmt.Println("=================================================RegisterFeeSplit========================================================")
	registerResponse, err := cli.Feesplit().RegisterFeeSplit(fromInfo, passWd, accInfo.GetAccountNumber(), seq, "register", "0x113a5369FAC959AFCeEB697981c39B5180813a7C", []uint64{75}, "")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(registerResponse)
	withdrawAddress := "fb1jyvuazl2nz0wd5v4zn62swg0x5pk38tefk3gcf"

	fmt.Println("=================================================UpdateFeeSplit========================================================")
	seq++
	updateResponse, err := cli.Feesplit().UpdateFeeSplit(fromInfo, passWd, accInfo.GetAccountNumber(), seq, "update", "0x113a5369FAC959AFCeEB697981c39B5180813a7C", withdrawAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(updateResponse)

	fmt.Println("=================================================QueryFeesplits========================================================")
	queryFeesplits, err := cli.Feesplit().QueryFeesplits(&query.PageRequest{
		Key:        nil,
		Offset:     0,
		Limit:      100,
		CountTotal: false,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(queryFeesplits)

	fmt.Println("=================================================QueryFeeSplit========================================================")
	queryFeesplit, err := cli.Feesplit().QueryFeeSplit("0x113a5369FAC959AFCeEB697981c39B5180813a7C")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(queryFeesplit)

	fmt.Println("=================================================QueryDeployerFeeSplits========================================================")
	queryDeployerFeeSplits, err := cli.Feesplit().QueryDeployerFeeSplits(fromInfo.GetAddress().String(), &query.PageRequest{
		Key:        nil,
		Offset:     0,
		Limit:      100,
		CountTotal: false,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(queryDeployerFeeSplits)

	fmt.Println("=================================================QueryWithdrawerFeeSplits========================================================")
	queryWithdrawerFeeSplits, err := cli.Feesplit().QueryWithdrawerFeeSplits("fb1jyvuazl2nz0wd5v4zn62swg0x5pk38tefk3gcf", &query.PageRequest{
		Key:        nil,
		Offset:     0,
		Limit:      100,
		CountTotal: false,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(queryWithdrawerFeeSplits)

	fmt.Println("=================================================QueryParams========================================================")
	queryParams, err := cli.Feesplit().QueryParams()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(queryParams)

	fmt.Println("=================================================CancelFeeSplit========================================================")
	seq++
	cancelResponse, err := cli.Feesplit().CancelFeeSplit(fromInfo, passWd, accInfo.GetAccountNumber(), seq, "cancel", "0x113a5369FAC959AFCeEB697981c39B5180813a7C")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cancelResponse)
}
