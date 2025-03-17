package aptos

import (
	"encoding/hex"

	"git.sonr.io/pkg/crypto/ed25519"
	"github.com/okx/go-wallet-sdk/wallet"
)

const HexPrefix = "0x"

type AptosWallet struct {
	wallet.WalletBase
}

func (aw *AptosWallet) GetRandomPrivateKey() (string, error) {
	p, err := ed25519.GenerateKey()
	if err != nil {
		return "", err
	}
	return HexPrefix + hex.EncodeToString(p), nil
}
