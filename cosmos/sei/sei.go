package sei

import (
	"github.com/sonr-io/coins/cosmos"
)

const (
	HRP = "sei"
)

func NewAddress(privateKeyHex string) (string, error) {
	return cosmos.NewAddress(privateKeyHex, HRP, false)
}

func ValidateAddress(address string) bool {
	return cosmos.ValidateAddress(address, HRP)
}
