module github.com/sonr-io/coins/filecoin

go 1.24.1

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.4
	github.com/dchest/blake2b v1.0.0
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0
	github.com/fxamacker/cbor v1.5.1
	github.com/sonr-io/crypto v0.3.0
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace launchpad.net/gocheck => gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
