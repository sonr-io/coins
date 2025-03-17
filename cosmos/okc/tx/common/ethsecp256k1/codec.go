package ethsecp256k1

import (
	"git.sonr.io/pkg/coins/cosmos/okc/tx/amino"
)

// RegisterCodec registers all the necessary types with amino for the given
// codec.
func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterConcrete(PubKey{}, PubKeyName, nil)
	cdc.RegisterConcrete(PrivKey{}, PrivKeyName, nil)
}

func init() {
	RegisterCodec(amino.GCodec)
}
