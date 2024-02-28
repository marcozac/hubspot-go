//go:build ignore

package main

import (
	"log"
	"text/template"

	"github.com/marcozac/hubspot-go/util"
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
	return util.WriteTemplateToFile(tmpl, "object_client.go", clients)
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
