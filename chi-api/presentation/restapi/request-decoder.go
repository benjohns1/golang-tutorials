package restapi

import (
	"io"
)

// RequestDecoder wraps a decoder factory function that creates a decoder to decode data from a request
type RequestDecoder struct {
	DecoderFactory func(r io.Reader) Decoder
}

// Decoder decodes data and stores in v
type Decoder interface {
	Decode(v interface{}) error
}

// Decode decodes data from the io.Reader and stores in v
func (d *RequestDecoder) Decode(r io.Reader, v interface{}) error {
	decoder := d.DecoderFactory(r)
	return decoder.Decode(&v)
}
