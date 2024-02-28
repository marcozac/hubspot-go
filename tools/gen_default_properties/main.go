//go:build ignore

// This program generates the default_properties.go file.
//
// WARNING
// This program requires a HubSpot Access Token with an high level of
// permissions. Although it doesn't modify any data, it's recommended to
// run it only in a development environment.
//
// The Access Token must be set in the HUBSPOT_TOKEN environment variable.
//
// It gets the default properties of the following objects:
//   - Contact
//   - Company
//   - Deal
//   - FeedbackSubmission
//   - LineItem
//   - Product
//   - Quote
//   - Discount
//   - Fee
//   - Tax
//   - Ticket
//   - Goal
//
// Custom properties are ignored.
package main

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/marcozac/hubspot-go"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error { //nolint:funlen
	tmpl, err := template.ParseFiles("template/default_properties.tmpl")
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ts, err := hubspot.NewEnvTokenSource("HUBSPOT_TOKEN")
	if err != nil {
		return err
	}
	client, err := hubspot.NewDefaultClient(ts, hubspot.WithContext(ctx))
	if err != nil {
		return err
	}
	dps := []*ObjectDefaultProperties{
		{
			Name: "ContactDefaultProperties",
			poc:  client.Properties.Contact,
		},
		{
			Name: "CompanyDefaultProperties",
			poc:  client.Properties.Company,
		},
		{
			Name: "DealDefaultProperties",
			poc:  client.Properties.Deal,
		},
		{
			Name: "FeedbackSubmissionDefaultProperties",
			poc:  client.Properties.FeedbackSubmission,
		},
		{
			Name: "LineItemDefaultProperties",
			poc:  client.Properties.LineItem,
		},
		{
			Name: "ProductDefaultProperties",
			poc:  client.Properties.Product,
		},
		{
			Name: "QuoteDefaultProperties",
			poc:  client.Properties.Quote,
		},
		{
			Name: "DiscountDefaultProperties",
			poc:  client.Properties.Discount,
		},
		{
			Name: "FeeDefaultProperties",
			poc:  client.Properties.Fee,
		},
		{
			Name: "TaxDefaultProperties",
			poc:  client.Properties.Tax,
		},
		{
			Name: "TicketDefaultProperties",
			poc:  client.Properties.Ticket,
		},
		{
			Name: "GoalDefaultProperties",
			poc:  client.Properties.Goal,
		},
	}
	for _, dp := range dps {
		ps, err := dp.poc.List(ctx)
		if err != nil {
			return err
		}
		dp.Fields = make([]Field, 0, len(ps))
		for _, p := range ps {
			// Ignore custom properties.
			if p.CreatedUserId != "" {
				continue
			}
			var f Field
			var sb strings.Builder
			ss := strings.Split(p.Name, "_")
			for _, s := range ss {
				if len(s) == 0 {
					// Keep the underscore if the string is empty.
					// It happens when the property name is, for example,
					// "hs__something". The underscore is kept to avoid
					// collisions with other properties.
					sb.WriteString("_")
					continue
				}
				sb.WriteString(strings.ToUpper(s[:1]))
				if len(s) > 1 {
					sb.WriteString(s[1:])
				}
			}
			f.Name = sb.String()
			f.Tag = fmt.Sprintf("`json:\"%s,omitempty\"`", p.Name)
			f.Description = p.Description
			switch p.Type {
			case hubspot.PropertyTypeString, hubspot.PropertyTypePhoneNumber:
				f.Type = "string"
			case hubspot.PropertyTypeNumber:
				// Set the type to Int64 for the contact hs_object_id property.
				// See: https://developers.hubspot.com/changelog/increasing-the-size-of-contact-record-ids
				if dp.Name == "ContactDefaultProperties" && p.Name == "hs_object_id" {
					f.Type = "Int64"
				} else {
					f.Type = "Int"
				}
			case hubspot.PropertyTypeDate:
				// must to be a pointer to respect the omitempty tag
				f.Type = "*Date"
			case hubspot.PropertyTypeDateTime:
				// must to be a pointer to respect the omitempty tag
				f.Type = "*DateTime"
			case hubspot.PropertyTypeEnumeration:
				f.Type = "Enumeration"
			case hubspot.PropertyTypeBool:
				f.Type = "Bool"
			default:
				return fmt.Errorf("unknown property type: %s", p.Type)
			}
			dp.Fields = append(dp.Fields, f)
		}
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, dps); err != nil {
		return err
	}
	data, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	f, err := os.Create("default_properties.go")
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}

type ObjectDefaultProperties struct {
	// Name is the name of the struct that will be generated.
	Name string

	// Fields is a list of fields that will be generated.
	Fields []Field

	poc *hubspot.PropertiesObjectClient
}

type Field struct {
	// Name is the name of the field.
	Name string

	// Type is the type of the field.
	Type string

	// Description is the description of the field.
	Description string

	// Tag is the json struct tag of the field.
	Tag string
}
