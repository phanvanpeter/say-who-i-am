package data

import (
	"encoding/json"
	"io"
)

// ToJSON deserializes an interface into a JSON object.
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON serializes a Reader (usually a Body of the Request) unto a given type (passed as interface).
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
