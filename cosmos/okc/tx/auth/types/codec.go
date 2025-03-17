package types

import (
	"git.sonr.io/pkg/coins/cosmos/okc/tx/amino"
)

// RegisterCodec registers concrete types on the codec
func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterConcrete(&StdTx{}, "cosmos-sdk/StdTx", nil)
}

func init() {
	RegisterCodec(amino.GCodec)
}
