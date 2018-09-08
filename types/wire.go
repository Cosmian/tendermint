package types

import (
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/encoding/amino"
)

type registerFunc = func(cdc *amino.Codec)

var registerFuncs []registerFunc

var cdc = amino.NewCodec()

// RegisterAminoDefaults registers with amino the types used by default with Commands
func RegisterAminoDefaults() {
	RegisterAmino(cryptoAmino.RegisterAmino)
	RegisterAmino(RegisterEvidences)
}

// RegisterAmino is used to register types with amino that are used by Commands
// see cryptoAmino.RegisterAmino() for an example registration
func RegisterAmino(register registerFunc) {
	registerFuncs = append(registerFuncs, register)
	register(cdc)
}

// RegisterBlockAmino registers on the passed Codec the registration made for types
func RegisterBlockAmino(cdc *amino.Codec) {
	for _, f := range registerFuncs {
		f(cdc)
	}
}
