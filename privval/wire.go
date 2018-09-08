package privval

import (
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

// RegisterAminoDefaults registers with amino the types used by default with private validator
func RegisterAminoDefaults() {
	RegisterAmino(cryptoAmino.RegisterAmino)
	RegisterAmino(RegisterSocketPVMsg)
}

// RegisterAmino is used to register types with amino that are used by private validator
// see cryptoAmino.RegisterAmino() for an example registration
func RegisterAmino(register func(cdc *amino.Codec) ) {
	register(cdc)
}
