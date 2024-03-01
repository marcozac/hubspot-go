package hsc

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/imports"

	"github.com/marcozac/hubspot-go"
)

func TestResolveImports(t *testing.T) {
	graph := &Graph{
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
	assert.NoError(t, graph.resolveImports())
}

func TestTemplate(t *testing.T) {
	tmpl, err := template.ParseFS(TemplateFS, "template/*.tmpl")
	require.NoError(t, err)

	graph := &Graph{
		PackageName: "hubspot",
		Header:      "Code generated by hsc. DO NOT EDIT.",
		Objects: []*Object{
			{
				Name: "Contact",
				Properties: []*Property{
					{
						Property: &hubspot.Property{
							Name:      "firstname",
							Label:     "First Name",
							Type:      hubspot.PropertyTypeString,
							FieldType: hubspot.PropertyFieldTypeText,
							GroupName: "contactinformation",
						},
					},
					{
						Property: &hubspot.Property{
							Name:      "lastname",
							Label:     "Last Name",
							Type:      hubspot.PropertyTypeString,
							FieldType: hubspot.PropertyFieldTypeText,
							GroupName: "contactinformation",
						},
					},
					{
						Property: &hubspot.Property{
							Name:        "my_date",
							Label:       "My Date",
							Type:        hubspot.PropertyTypeDate,
							FieldType:   hubspot.PropertyFieldTypeDate,
							GroupName:   "contactinformation",
							Description: "A custom date field",
						},
					},
				},
				EndpointTarget: &EndpointTarget{
					Name:    "Contacts",
					Package: "github.com/marcozac/hubspot-go/endpoint",
				},
			},
			{
				Name: "ContactAlias",
				Properties: []*Property{
					{
						Property: &hubspot.Property{
							Name:        "firstname",
							Label:       "First Name",
							Type:        hubspot.PropertyTypeString,
							FieldType:   hubspot.PropertyFieldTypeText,
							GroupName:   "contactinformation",
							Description: "The first name of the contact",
						},
					},
					{
						Property: &hubspot.Property{
							Name:        "lastname",
							Label:       "Last Name",
							Type:        hubspot.PropertyTypeString,
							FieldType:   hubspot.PropertyFieldTypeText,
							GroupName:   "contactinformation",
							Description: "The last name of the contact",
						},
					},
				},
				EndpointTarget: &EndpointTarget{
					Name:    "Example",
					Package: "github.com/marcozac/hubspot-go/hsc/internal/endpoint-alias",
				},
			},
		},
	}
	assert.NoError(t, graph.resolveImports())

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, "objects_client", graph)
	require.NoError(t, err)

	ff, err := imports.Process("", buf.Bytes(), nil)

	assert.NoError(t, err)
	t.Log(string(ff))
}
