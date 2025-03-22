module github.com/sonr-io/coins/helium

go 1.24.1

require (
	github.com/golang/protobuf v1.5.4
	github.com/shopspring/decimal v1.3.1
	github.com/sonr-io/crypto v0.3.0
	github.com/stretchr/testify v1.10.0
	golang.org/x/crypto v0.36.0
	google.golang.org/protobuf v1.36.5
)

require (
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace launchpad.net/gocheck => gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
