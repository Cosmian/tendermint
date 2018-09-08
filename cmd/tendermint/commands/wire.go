package commands

import (
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

// RegisterAminoDefaults registers with amino the types used by default with Commands
func RegisterAminoDefaults() {
	RegisterAmino(cryptoAmino.RegisterAmino)
}

// RegisterAmino is used to register types with amino that are used by Commands
// see cryptoAmino.RegisterAmino() for an example registration
func RegisterAmino(register func(cdc *amino.Codec) ) {
	register(cdc)
}
