module github.com/sonr-io/coins

go 1.24.2

require (
	github.com/btcsuite/btcd/btcutil v1.1.6
	github.com/sonr-io/coins/aptos v0.0.4
	github.com/sonr-io/coins/cosmos v0.0.4
)

replace github.com/sonr-io/coins/aptos => ./aptos
replace github.com/sonr-io/coins/cosmos => ./cosmos

require (
	github.com/btcsuite/btcd v0.24.2 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.4 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0 // indirect
	github.com/ethereum/go-ethereum v1.15.5 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/okx/go-wallet-sdk/wallet v0.0.0-20231109124131-23d8b0dd4b6f // indirect
	github.com/sonr-io/crypto v0.3.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
)
