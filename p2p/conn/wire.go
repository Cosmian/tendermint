package conn

import (
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/encoding/amino"
)

var cdc *amino.Codec = amino.NewCodec()

// RegisterAminoDefaults registers with amino the types used by default with p2p/conn
func RegisterAminoDefaults() {
	RegisterAmino(cryptoAmino.RegisterAmino)
	RegisterAmino(RegisterPacket)
}

// RegisterAmino is used to register types with amino that are used by p2p/conn
// see cryptoAmino.RegisterAmino() for an example registration
func RegisterAmino(register func(cdc *amino.Codec)) {
	register(cdc)
}
