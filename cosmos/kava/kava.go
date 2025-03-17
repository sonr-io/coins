package kava

import (
	"git.sonr.io/pkg/coins/cosmos"
)

const (
	HRP = "kava"
)

func NewAddress(privateKeyHex string) (string, error) {
	return cosmos.NewAddress(privateKeyHex, HRP, false)
}

func ValidateAddress(address string) bool {
	return cosmos.ValidateAddress(address, HRP)
}
