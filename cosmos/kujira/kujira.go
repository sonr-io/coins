package kujira

import "github.com/sonr-io/coins/cosmos"

const (
	HRP = "kujira"
)

func NewAddress(privateKeyHex string) (string, error) {
	return cosmos.NewAddress(privateKeyHex, HRP, false)
}

func ValidateAddress(address string) bool {
	return cosmos.ValidateAddress(address, HRP)
}
