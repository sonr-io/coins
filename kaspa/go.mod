module git.sonr.io/pkg/coins/kaspa

go 1.23.0

toolchain go1.24.1

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.4
	github.com/kaspanet/go-secp256k1 v0.0.7
	github.com/kaspanet/kaspad v0.12.14
	github.com/stretchr/testify v1.10.0
	golang.org/x/crypto v0.36.0
)

require (
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/decred/dcrd/crypto/blake256 v1.1.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/jrick/logrotate v1.0.0 // indirect
	github.com/kaspanet/go-muhash v0.0.4 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7 // indirect
	golang.org/x/sys v0.31.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace launchpad.net/gocheck => gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
