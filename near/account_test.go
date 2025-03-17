package near

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"testing"

	"git.sonr.io/pkg/crypto/base58"
	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	t.Log("privateKey : ", hex.EncodeToString(privateKey))
	address := base58.Encode(publicKey)
	t.Log("address : ", address)
}

func TestValidateAddress(t *testing.T) {
	addr, _, err := NewAccount()
	require.NoError(t, err)
	ret := ValidateAddress(addr)
	require.True(t, ret)
}
