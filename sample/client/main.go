package main

import (
	"log"

	gosdk "github.com/metaofo/fbchain-sdk"
	"github.com/metaofo/fbchain-sdk/utils"
)

const (
	// TODO: link to mainnet of FbChain later
	rpcURL = "tcp://8.210.109.185:26657"
	// user's name
	name = "alice"
	// user's mnemonic
	mnemonic = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	// user's password
	passWd = "12345678"
	// target address
	addr     = "fb1qj5c07sm6jetjz8f509qtrxgh4psxkv3htfcuj"
	baseCoin = "fibo"
)

func main() {
	//-------------------- 1. preparation --------------------//
	// NOTE: either of the both ways below to pay fees is available

	// WAY 1: create a client config with fixed fees
	config, err := gosdk.NewClientConfig(rpcURL, "fbc-1230", gosdk.BroadcastBlock, "0.01fibo", 200000,
		0, "")
	if err != nil {
		log.Fatal(err)
	}

	// WAY 2: alternative client config with the fees by auto gas calculation
	config, err = gosdk.NewClientConfig(rpcURL, "fbc-1230", gosdk.BroadcastBlock, "", 200000,
		1.1, "0.000000001fibo")
	if err != nil {
		log.Fatal(err)
	}

	cli := gosdk.NewClient(config)

	// create an account with your own mnemonic，name and password
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	if err != nil {
		log.Fatal(err)
	}

	fbAddress := fromInfo.GetAddress().String()
	log.Println(fbAddress)
	//-------------------- 2. query for the information of your address --------------------//
	accInfo, err := cli.Auth().QueryAccount(fbAddress)
	if err != nil {
		log.Fatal(err)
	}

	coins := accInfo.GetCoins()

	log.Println("address: ", accInfo.GetAddress().String(), ", coins: ", coins)

	//-------------------- 3. transfer to other address --------------------//

	// sequence number of the account must be increased by 1 whenever a transaction of the account takes effect
	accountNum, sequenceNum := accInfo.GetAccountNumber(), accInfo.GetSequence()
	res, err := cli.Token().Send(fromInfo, passWd, addr, "1"+baseCoin, "my memo", accountNum, sequenceNum)
	if err != nil {
		log.Fatal("send tx: ", err)
	}

	log.Println(res)

	//-------------------- 4. deposit for staking --------------------//

	// increase sequence number
	sequenceNum++
	res, err = cli.Staking().Deposit(fromInfo, passWd, "0.1"+baseCoin, "my memo", accountNum, sequenceNum)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
