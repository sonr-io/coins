package coins

import (
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/sonr-io/coins/aptos"
	"github.com/sonr-io/coins/cosmos/atom"
	"github.com/sonr-io/coins/cosmos/axelar"
	"github.com/sonr-io/coins/cosmos/cronos"
	"github.com/sonr-io/coins/cosmos/evmos"
	"github.com/sonr-io/coins/cosmos/iris"
	"github.com/sonr-io/coins/cosmos/juno"
	"github.com/sonr-io/coins/cosmos/kava"
	"github.com/sonr-io/coins/cosmos/kujira"
	"github.com/sonr-io/coins/cosmos/okc"
	"github.com/sonr-io/coins/cosmos/osmo"
	"github.com/sonr-io/coins/cosmos/secret"
	"github.com/sonr-io/coins/cosmos/sei"
	"github.com/sonr-io/coins/cosmos/stargaze"
	"github.com/sonr-io/coins/cosmos/tia"
)

var errUnimplemented = errors.New("unimplemented")

// AptosAddressFromPublicKey returns the Aptos address from the public key
func AptosAddressFromPublicKey(publicKey []byte) (string, error) {
	return aptos.NewPubKeyAddress(hex.EncodeToString(publicKey), true)
}

// AtomAddressFromPublicKey returns the Atom address from the public key
func AtomAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, atom.HRP)
}

// AxelarAddressFromPublicKey returns the Axelar address from the public key
func AxelarAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, axelar.HRP)
}

// AvaxAddressFromPublicKey returns the Avax address from the public key
func AvaxAddressFromPublicKey(publicKey []byte) (string, error) {
	return "", errUnimplemented
}

// BitcoinAddressFromPublicKey returns the Bitcoin address from the public key
func BitcoinAddressFromPublicKey(publicKey []byte) (string, error) {
	return "", errUnimplemented
}

// CosmosAddressFromPublicKey returns the Cosmos address from the public key
func CosmosAddressFromPublicKey(publicKey []byte, hrp string) (string, error) {
	bytes := btcutil.Hash160(publicKey)
	return bech32.EncodeFromBase256(hrp, bytes)
}

// CronosAddressFromPublicKey returns the Cronos address from the public key
func CronosAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, cronos.HRP)
}

// EthereumAddressFromPublicKey returns the Ethereum address from the public key
func EthereumAddressFromPublicKey(publicKey []byte) (string, error) {
	return "", errUnimplemented
}

// EvmosAddressFromPublicKey returns the Evmos address from the public key
func EvmosAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, evmos.HRP)
}

// FilecoinAddressFromPublicKey returns the Filecoin address from the public key
func FilecoinAddressFromPublicKey(publicKey []byte) (string, error) {
	return "", errUnimplemented
}

// HeliumAddressFromPublicKey returns the Helium address from the publicKey
func HeliumAddressFromPublicKey(publicKey []byte) (string, error) {
	return "", errUnimplemented
}

// IrisAddressFromPublicKey returns the Iris address from the public key
func IrisAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, iris.HRP)
}

// JunoAddressFromPublicKey returns the Juno address from the public key
func JunoAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, juno.HRP)
}

// KavaAddressFromPublicKey returns the Kava address from the public key
func KavaAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, kava.HRP)
}

// KujiraAddressFromPublicKey returns the Kujira address from the public key
func KujiraAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, kujira.HRP)
}

// OkcAddressFromPublicKey returns the Okc address from the public key
func OkcAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, okc.HRP)
}

// OsmoAddressFromPublicKey returns the Osmo address from the public key
func OsmoAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, osmo.HRP)
}

// PolkadotAddressFromPublicKey returns the Polkadot address from the public key
func PolkadotAddressFromPublicKey(publicKey []byte) (string, error) {
	return "", errUnimplemented
}

// SecretAddressFromPublicKey returns the Secret address from the public key
func SecretAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, secret.HRP)
}

// SeiAddressFromPublicKey returns the Sei address from the public key
func SeiAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, sei.HRP)
}

// SolanaAddressFromPublicKey returns the Solana address from the public key
func SolanaAddressFromPublicKey(publicKey []byte) (string, error) {
	return "", errUnimplemented
}

// SolanaAddressFromPublicKey returns the Solana address from the public key
func SonrAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, "idx")
}

// StargazeAddressFromPublicKey returns the Stargaze address from the public key
func StargazeAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, stargaze.HRP)
}

// SuiAddressFromPublicKey returns the Sui address from the public key
func SuiAddressFromPublicKey(publicKey []byte) (string, error) {
	return "", errUnimplemented
}

// TezosAddressFromPublicKey returns the Tezos address from the public key
func TezosAddressFromPublicKey(publicKey []byte) (string, error) {
	return "", errUnimplemented
}

// TiaAddressFromPublicKey returns the Tia address from the public key
func TiaAddressFromPublicKey(publicKey []byte) (string, error) {
	return CosmosAddressFromPublicKey(publicKey, tia.HRP)
}
