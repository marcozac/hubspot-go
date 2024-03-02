package main

import (
	"context"
	"fmt"
	"log"

	"github.com/marcozac/hubspot-go"
	"github.com/marcozac/hubspot-go/hsc"
	"github.com/marcozac/hubspot-go/limiter"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error { //nolint:funlen
	const endpointPkg = "github.com/marcozac/hubspot-go/endpoint"
	g := &hsc.Graph{
		PackagePath: "github.com/marcozac/hubspot-go",
		Objects: []*hsc.Object{
			{
				Name: "Contact",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "Contacts",
					Package: endpointPkg,
				},
			},
			{
				Name: "Company",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "Companies",
					Package: endpointPkg,
				},
			},
			{
				Name: "Deal",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "Deals",
					Package: endpointPkg,
				},
			},
			{
				Name: "FeedbackSubmission",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "FeedbackSubmissions",
					Package: endpointPkg,
				},
			},
			{
				Name: "LineItem",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "LineItems",
					Package: endpointPkg,
				},
			},
			{
				Name: "Product",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "Products",
					Package: endpointPkg,
				},
			},
			{
				Name: "Quote",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "Quotes",
					Package: endpointPkg,
				},
			},
			{
				Name: "Discount",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "Discounts",
					Package: endpointPkg,
				},
			},
			{
				Name: "Fee",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "Fees",
					Package: endpointPkg,
				},
			},
			{
				Name: "Tax",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "Taxes",
					Package: endpointPkg,
				},
			},
			{
				Name: "Ticket",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "Tickets",
					Package: endpointPkg,
				},
			},
			{
				Name: "Goal",
				EndpointTarget: &hsc.EndpointTarget{
					Name:    "Goals",
					Package: endpointPkg,
				},
			},
		},
	}
	ts, err := hubspot.NewEnvTokenSource("HUBSPOT_TOKEN")
	if err == nil {
		ctx := context.Background()
		client, err := hubspot.NewDefaultClient(ts, hubspot.WithLimiter(limiter.NewDefault10()))
		if err != nil {
			return fmt.Errorf("create default client: %w", err)
		}
		for _, o := range g.Objects {
			var pc *hubspot.PropertiesObjectClient
			switch o.Name {
			case "Contact":
				pc = client.Properties.Contact
			case "Company":
				pc = client.Properties.Company
			case "Deal":
				pc = client.Properties.Deal
			case "FeedbackSubmission":
				pc = client.Properties.FeedbackSubmission
			case "LineItem":
				pc = client.Properties.LineItem
			case "Product":
				pc = client.Properties.Product
			case "Quote":
				pc = client.Properties.Quote
			case "Discount":
				pc = client.Properties.Discount
			case "Fee":
				pc = client.Properties.Fee
			case "Tax":
				pc = client.Properties.Tax
			case "Ticket":
				pc = client.Properties.Ticket
			case "Goal":
				pc = client.Properties.Goal
			default:
				return fmt.Errorf("unexpected object: %s", o.Name)
			}
			props, err := pc.List(ctx)
			if err != nil {
				return fmt.Errorf("list %s properties: %w", o.Name, err)
			}
			o.Properties = make([]*hsc.Property, 0, len(props))
			for _, p := range props {
				o.Properties = append(o.Properties, &hsc.Property{
					Property: p,
				})
			}
		}
		if err := g.Init(); err != nil {
			return fmt.Errorf("init graph: %w", err)
		}
		if err := hsc.GenerateObjectsTypeDefault(g); err != nil {
			return fmt.Errorf("generate objects type default: %w", err)
		}
	} else {
		fmt.Println("Failed to get token from environment, skipping default properties generation")
		if err := g.Init(); err != nil {
			return fmt.Errorf("init graph: %w", err)
		}
	}
	if err := hsc.GenerateObjectsClient(g); err != nil {
		return fmt.Errorf("generate objects client: %w", err)
	}
	return nil
}
