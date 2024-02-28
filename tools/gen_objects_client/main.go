//go:build ignore

package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"text/template"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	tmpl, err := template.ParseFiles("template/object_client.tmpl")
	if err != nil {
		return err
	}
	clients := []ObjectClient{
		{
			Name:           "Contact",
			EndpointTarget: "Contacts",
		},
		{
			Name:           "Company",
			EndpointTarget: "Companies",
		},
		{
			Name:           "Deal",
			EndpointTarget: "Deals",
		},
		{
			Name:           "FeedbackSubmission",
			EndpointTarget: "FeedbackSubmissions",
		},
		{
			Name:           "LineItem",
			EndpointTarget: "LineItems",
		},
		{
			Name:           "Product",
			EndpointTarget: "Products",
		},
		{
			Name:           "Quote",
			EndpointTarget: "Quotes",
		},
		{
			Name:           "Discount",
			EndpointTarget: "Discounts",
		},
		{
			Name:           "Fee",
			EndpointTarget: "Fees",
		},
		{
			Name:           "Tax",
			EndpointTarget: "Taxes",
		},
		{
			Name:           "Ticket",
			EndpointTarget: "Tickets",
		},
		{
			Name:           "Goal",
			EndpointTarget: "Goals",
		},
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, clients); err != nil {
		return err
	}
	data, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	f, err := os.Create("object_client.go")
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}

type ObjectClient struct {
	// Name is the name of the object type that the client is for.
	// It should be related to the endpoint name.
	//
	// Example:
	// 	"Contact" => endpoint.Contacts
	Name string

	// EndpointTarget is the target endpoint for the client.
	// It must be equal to an exported constant in the endpoint package.
	//
	// Example:
	// 	"Contacts" => endpoint.Contacts
	EndpointTarget string
}
