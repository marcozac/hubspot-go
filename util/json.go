package util

func MarshalStringerAsJSON(s Stringer) ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}
