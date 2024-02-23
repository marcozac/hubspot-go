package property

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

type Bool bool

func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}

var _ json.Marshaler = Bool(true)

func (b Bool) MarshalJSON() ([]byte, error) {
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

type Enumeration []string

func (e Enumeration) String() string {
	return strings.Join(e, ";")
}

var _ json.Marshaler = Enumeration{}

func (e Enumeration) MarshalJSON() ([]byte, error) {
	return []byte(`"` + e.String() + `"`), nil
}

var _ json.Unmarshaler = (*Enumeration)(nil)

func (e *Enumeration) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	for _, v := range strings.Split(s, ";") {
		*e = append(*e, v)
	}
	return nil
}
