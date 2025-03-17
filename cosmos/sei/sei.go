package sei

import (
	"git.sonr.io/pkg/coins/cosmos"
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
