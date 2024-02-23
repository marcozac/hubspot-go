package property

import (
	"bytes"
	"encoding/json"
)

type Bool bool

func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}

var _ json.Marshaler = (*Bool)(nil)

func (b *Bool) MarshalJSON() ([]byte, error) {
	return []byte(`"` + b.String() + `"`), nil
}

var (
	_ json.Unmarshaler = (*Bool)(nil)

	trueStringBytes  = []byte(`"true"`)
	falseStringBytes = []byte(`"false"`)
)

func (b *Bool) UnmarshalJSON(data []byte) error {
	switch {
	case bytes.Equal(data, trueStringBytes):
		*b = true
	case bytes.Equal(data, falseStringBytes):
		*b = false
	default:
		// Try to unmarshal as a bool.
		return json.Unmarshal(data, (*bool)(b))
	}
	return nil
}
