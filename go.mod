module github.com/sonr-io/coins

go 1.24.2

require (
	github.com/btcsuite/btcd/btcutil v1.1.6
	github.com/sonr-io/coins/aptos v0.0.5
	github.com/sonr-io/coins/bitcoin v0.0.5
	github.com/sonr-io/coins/cosmos v0.0.5
	github.com/sonr-io/coins/ethereum v0.0.5
	github.com/sonr-io/coins/filecoin v0.0.0-00010101000000-000000000000
	github.com/sonr-io/coins/solana v0.0.5
)

replace github.com/sonr-io/coins/aptos => ./aptos

replace github.com/sonr-io/coins/bitcoin => ./bitcoin

replace github.com/sonr-io/coins/cosmos => ./cosmos

replace github.com/sonr-io/coins/ethereum => ./ethereum

replace github.com/sonr-io/coins/filecoin => ./filecoin

replace github.com/sonr-io/coins/helium => ./helium

replace github.com/sonr-io/coins/polkadot => ./polkadot

replace github.com/sonr-io/coins/solana => ./solana

replace github.com/sonr-io/coins/sui => ./sui

replace github.com/sonr-io/coins/tezos => ./tezos

require (
	github.com/bits-and-blooms/bitset v1.20.0 // indirect
	github.com/btcsuite/btcd v0.24.2 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.4 // indirect
	github.com/btcsuite/btcd/btcutil/psbt v1.1.8 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce // indirect
	github.com/consensys/bavard v0.1.27 // indirect
	github.com/consensys/gnark-crypto v0.16.0 // indirect
	github.com/crate-crypto/go-ipa v0.0.0-20240724233137-53bbb0ceb27a // indirect
	github.com/crate-crypto/go-kzg-4844 v1.1.0 // indirect
	github.com/dchest/blake2b v1.0.0 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.1.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0 // indirect
	github.com/ethereum/c-kzg-4844 v1.0.0 // indirect
	github.com/ethereum/go-ethereum v1.15.5 // indirect
	github.com/ethereum/go-verkle v0.2.2 // indirect
	github.com/fxamacker/cbor v1.5.1 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/okx/go-wallet-sdk/wallet v0.0.0-20231109124131-23d8b0dd4b6f // indirect
	github.com/sonr-io/crypto v0.3.0 // indirect
	github.com/supranational/blst v0.3.14 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
