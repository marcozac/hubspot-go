package property

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBool(t *testing.T) {
	assert.Equal(t, "true", Bool(true).String(), "Bool.String() should return 'true' for true")
	assert.Equal(t, "false", Bool(false).String(), "Bool.String() should return 'false' for false")
	type S struct {
		B Bool `json:"b"`
	}
	t.Run("Marshal", func(t *testing.T) {
		v := &S{B: true}
		data, err := json.Marshal(v)
		require.NoError(t, err, "json.Marshal must not return an error")
		assert.Equal(t, []byte(`{"b":"true"}`), data)
	})
	t.Run("Unmarshal", func(t *testing.T) {
		var s S
		assert.NoError(t, json.Unmarshal([]byte(`{ "b": "true" }`), &s), `json.Unmarshal: "true"`)
		assert.True(t, bool(s.B))
		assert.NoError(t, json.Unmarshal([]byte(`{ "b": "false" }`), &s), `json.Unmarshal: "false"`)
		assert.False(t, bool(s.B))
		assert.NoError(t, json.Unmarshal([]byte(`{ "b": true }`), &s), `json.Unmarshal: true`)
		assert.True(t, bool(s.B))
		assert.NoError(t, json.Unmarshal([]byte(`{ "b": false }`), &s), `json.Unmarshal: true`)
		assert.False(t, bool(s.B))
		assert.Error(t, json.Unmarshal([]byte(`{ "b": "foo" }`), &s), `json.Unmarshal: "foo"`)
	})
}
