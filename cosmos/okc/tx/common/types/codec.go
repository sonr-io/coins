package types

import (
	"github.com/sonr-io/coins/cosmos/okc/tx/amino"
)

// Register the sdk message type
func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterInterface((*Msg)(nil), nil)
	cdc.RegisterInterface((*Tx)(nil), nil)
}

func init() {
	RegisterCodec(amino.GCodec)
}
