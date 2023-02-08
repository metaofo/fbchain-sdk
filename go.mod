module github.com/metaofo/fbchain-sdk

go 1.16

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/ethereum/go-ethereum v1.10.8
	github.com/golang/mock v1.6.0
	github.com/stretchr/testify v1.8.0
	github.com/zhengjianfeng1103/fbc v0.0.0-20230208063344-8b5f6350b8d2
)

replace (
	github.com/ethereum/go-ethereum => github.com/okex/go-ethereum v1.10.8-oec3
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/go-amino => github.com/okex/go-amino v0.15.1-okc4
)
