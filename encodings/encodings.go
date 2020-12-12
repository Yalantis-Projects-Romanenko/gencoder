package encodings

import (
	"github.com/fdistorted/gencrypt/encodings/numbers"
	"github.com/fdistorted/gencrypt/encodings/piglatin"
)

var EncodeFunctions = map[string]func(string) string{
	"numbers":  numbers.Encode,
	"piglatin": piglatin.Encode,
}

var DecodeFunctions = map[string]func(string) string{
	"numbers": numbers.Decode,
}
