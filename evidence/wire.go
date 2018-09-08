package evidence

import (
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	"github.com/tendermint/tendermint/types"
)

var cdc = amino.NewCodec()

// RegisterAminoDefaults registers with amino the types used by the default with Evidence
func RegisterAminoDefaults() {
	RegisterAmino(RegisterEvidenceMessages)
	RegisterAmino(cryptoAmino.RegisterAmino)
	RegisterAmino(types.RegisterEvidences)
}

// RegisterAmino is used to register types with amino that are used by Evidence
// see cryptoAmino.RegisterAmino() for an example registration
func RegisterAmino(register func(cdc *amino.Codec) ) {
	register(cdc)
}

