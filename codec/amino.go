package codec

import (
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/encoding/amino"
)

// amino codec to marshal/unmarshal
type Amino struct {
	*amino.Codec
}

func New() *Amino {
	cdc := amino.NewCodec()
	return &Amino{cdc}
}

// Register the go-crypto to the codec
func RegisterCrypto(cdc *Amino) {
	cryptoAmino.RegisterAmino(cdc.Codec)
}

func (cdc *Amino) Cache() Codec {
	return newCache(cdc)
}

func (cdc *Amino) Seal() Codec {
	cdc.Codec.Seal()
	return cdc
}

//__________________________________________________________________

// generic sealed codec to be used throughout sdk
var Cdc Codec

func init() {
	cdc := New()
	RegisterCrypto(cdc)
	Cdc = cdc.Seal()
}