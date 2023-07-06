package main

import (
	"html/template"
	"os"
)

var (
	tmplt *template.Template
)

type TemplateData struct {
	Ingresses  []IngressesData
	Namespaces []string
}

type IngressesData struct {
	Name        string
	Link        string
	Namespace   string
	Annotations map[string]string
}

// Sample ingress data
var sampleData = TemplateData{
	Ingresses: []IngressesData{
		{
			Name:      "frontend-ingress",
			Link:      "https://prod.frontend.acme.com",
			Namespace: "production",
			Annotations: map[string]string{
				"acme.com/repository": "acme/frontend",
				"acme.com/contact":    "John Doe",
			},
		},
		{
			Name:      "backend-ingress",
			Link:      "https://prod.backend.acme.com",
			Namespace: "production",
			Annotations: map[string]string{
				"acme.com/repository": "acme/backend",
				"acme.com/contact":    "John Doe",
			},
		},
		{
			Name:      "frontend-ingress",
			Link:      "https://dev.frontend.acme.com",
			Namespace: "development",
			Annotations: map[string]string{
				"acme.com/repository": "acme/frontend",
				"acme.com/contact":    "John Doe",
			},
		},
		{
			Name:      "backend-ingress",
			Link:      "https://dev.backend.acme.com",
			Namespace: "development",
			Annotations: map[string]string{
				"acme.com/repository": "acme/backend",
				"acme.com/contact":    "John Doe",
			},
		},
		{
			Name:      "frontend-ingress",
			Link:      "https://staging.frontend.acme.com",
			Namespace: "staging",
			Annotations: map[string]string{
				"acme.com/repository": "acme/frontend",
				"acme.com/contact":    "John Doe",
			},
		},
		{
			Name:      "backend-ingress",
			Link:      "https://staging.backend.acme.com",
			Namespace: "staging",
			Annotations: map[string]string{
				"acme.com/repository": "acme/backend",
				"acme.com/contact":    "John Doe",
			},
		},
		{
			Name:      "frontend-ingress",
			Link:      "https://test.frontend.acme.com",
			Namespace: "testing",
			Annotations: map[string]string{
				"acme.com/repository": "acme/frontend",
				"acme.com/contact":    "Tim Tester",
			},
		},
		{
			Name:      "backend-ingress",
			Link:      "https://test.backend.acme.com",
			Namespace: "testing",
			Annotations: map[string]string{
				"acme.com/repository": "acme/backend",
				"acme.com/contact":    "Tim Tester",
			},
		},
	},
	Namespaces: []string{
		"production",
		"development",
		"staging",
		"testing",
	},
}

func main() {
	// Create output path
	f, _ := os.Create("index.html")

	// Parse template
	tmplt, _ = template.ParseFiles("site.html")

	// Write output to file
	tmplt.Execute(f, sampleData)

	// Close file
	f.Close()
}
