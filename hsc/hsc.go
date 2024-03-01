package hsc

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/tools/imports"
)

//go:embed template/*
var TemplateFS embed.FS

var tmpl = template.Must(template.ParseFS(TemplateFS, "template/*.tmpl"))

func GenreateObjectsType(g *Graph) error {
	return generate(g, "objects_type")
}

func GenerateObjectsTypeDefault(g *Graph) error {
	return generate(g, "objects_type_default")
}

func GenerateObjectsClient(g *Graph) error {
	return generate(g, "objects_client")
}

// generate generates the file with the given name.
//
// The Graph must be initialized before calling this function.
func generate(g *Graph, name string) error {
	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, name, g); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}
	data := buf.Bytes()
	fb, err := imports.Process("", data, nil)
	if err != nil {
		return fmt.Errorf("format file: %w", err)
	}
	filename := filepath.Join(g.OutDir, name+".go")
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer f.Close()
	if _, err := f.Write(fb); err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}
