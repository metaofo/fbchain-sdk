package evm

import (
	authcli "github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth/client/utils"
	"github.com/FiboChain/fbc/libs/tendermint/crypto/etherhash"
	evmtypes "github.com/FiboChain/fbc/x/evm/types"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core/types"
)

// GetTxHash calculates the tx hash
func (ec evmClient) GetTxHash(signedTx *ethcore.Transaction) (txHash ethcmn.Hash, err error) {
	v, r, s := signedTx.RawSignatureValues()
	tx := evmtypes.MsgEthereumTx{
		Data: evmtypes.TxData{
			AccountNonce: signedTx.Nonce(),
			Price:        signedTx.GasPrice(),
			GasLimit:     signedTx.Gas(),
			Recipient:    signedTx.To(),
			Amount:       signedTx.Value(),
			Payload:      signedTx.Data(),
			V:            v,
			R:            r,
			S:            s,
		},
	}

	txBytes, err := authcli.GetTxEncoder(ec.GetCodec(), authcli.WithEthereumTx())(&tx)
	if err != nil {
		return
	}

	txHash = ethcmn.BytesToHash(etherhash.Sum(txBytes))
	return
}
