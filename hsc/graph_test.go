package hsc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveImports(t *testing.T) {
	node := &Node{
		Objects: []*Object{
			{
				EndpointTarget: &EndpointTarget{
					Package: "github.com/marcozac/hubspot-go/endpoint",
				},
			},
			{
				EndpointTarget: &EndpointTarget{
					Package: "github.com/marcozac/hubspot-go/hsc/internal/endpoint-alias",
				},
			},
		},
	}
	assert.NoError(t, node.resolveImports())
}
