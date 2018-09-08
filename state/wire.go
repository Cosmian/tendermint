package state

import (
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

// RegisterAminoDefaults registers with amino the types used by the State
func RegisterAminoDefaults() {
	RegisterAmino(cryptoAmino.RegisterAmino)
}

// RegisterAmino is used to register types with amino that used by the state package
// see cryptoAmino.RegisterAmino() for an example registration
func RegisterAmino(register func(cdc *amino.Codec) ) {
	register(cdc)
}
