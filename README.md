## fbc Go SDK


fbc Go SDK is a lightweight Go library to interact with fbc.

### 1. Components

- client.go - The main client of GO SDK is created in this file. Developers are supposed to set up the config with own requirement during the client creation.
- expose - Abstraction with the interfaces of each module. The implements of it are filled in the folder `module`.
- module - The main logic for GO SDK queries and txs are classified by their own module names in fbc. Developers can find out the concrete designs under the specific module folder. Please focus on the files, tx.go and query.go. 
- mocks - Mock client tools for unit test of the main client in GO SDK.
- sample - A clear short user guide.
- types - The necessary struct set of fbc is built here. Developers are allowed to import some basic types like Dec and AccAddress directly if they want.
- utils -  A useful tool set for the one who is going to send more transactions and queries, is spilt by module names as the file names. Beyond, the operation of account keys with mnemonics remains in the file `account.go`.

### 2. Installation

Go version above 1.15 is required.

The developer can get the fbc Go SDK directly by `git clone` from github : https://github.com/okex/fbc-go-sdk

### 3. API

The api functions of transactions and queries are all under the path `expose`. You can find more details in fbc-docs : https://fbc-docs.readthedocs.io/en/latest/api/sdk/go-sdk.html

### 4. Tendermint query

fbc Go SDK also provides node query functions so that the underlying information of the blockchain is available for developers.

The tendermint query functions could be found in the file `exposed/tendermint.go `. Developers could make it through with the file `module/tendermint/query.go` and get clear how to invoke them.

### 5. Example

`Client` seems necessary to every operation with Go SDK. Here are the examples :

```go
	// rpcURL should be modified according to the actual situation
	rpcURL   = "https://rpc.fibochain.org"
	name     = "alice"
	passWd   = "12345678"
	mnemonic = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	addr     = "fb1qj5c07sm6jetjz8f509qtrxgh4psxkv3htfcuj"
	
	// build the client with own config
	config, _ := sdk.NewClientConfig(rpcURL, "fbc-1230", sdk.BroadcastBlock, "0.00002fibo", 200000, 0, "")
	client := sdk.NewClient(config)

	// create your account key info by 'name','passWd' and 'mnemonic'
	keyInfo, _, _ := utils.CreateAccountWithMnemo(mnemonic, name, passWd)

	// get info of your account from fbc
	accInfo, _ := client.Auth().QueryAccount(keyInfo.GetAddress().String())

	// transfer some fibo to addr
	res, _ := client.Token().Send(keyInfo, passWd, addr, "0.1024fibo", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())

```

You can invoke more and more api functions with the object `client`.

### 6. Testing

All changes and addition of codes will be pushed with unit tests strictly. 

### 7. Contributing

No doubt that it's admirable to make contributions to fbc Go SDK. You can provide your code as long as you have tested it with a local client and your unit test showed its validity.  

