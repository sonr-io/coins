module git.sonr.io/pkg/coins/flow

go 1.24.1

require (
	git.sonr.io/pkg/crypto/util v0.0.2
	github.com/btcsuite/btcd/btcec/v2 v2.3.4
	github.com/ethereum/go-ethereum v1.15.5
	github.com/stretchr/testify v1.10.0
	golang.org/x/crypto v0.36.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	golang.org/x/sys v0.31.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace launchpad.net/gocheck => gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
